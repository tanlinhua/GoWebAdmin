package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

// https://github.com/dvwright/xss-mw

type XssMw struct {
	FieldsToSkip []string
	BmPolicy     string
}

type XssMwJson map[string]interface{}

func (mw *XssMw) RemoveXss() gin.HandlerFunc {
	mw.FieldsToSkip = append(mw.FieldsToSkip, "password")
	return func(c *gin.Context) {
		mw.callRemoveXss(c)
	}
}

func (mw *XssMw) callRemoveXss(c *gin.Context) {
	if mw.BmPolicy == "" {
		mw.BmPolicy = "StrictPolicy"
	} else if mw.BmPolicy != "StrictPolicy" && mw.BmPolicy != "UGCPolicy" {
		fmt.Println("BlueMondy Policy setting is incorrect!")
		c.Abort()
		return
	}

	err := mw.XssRemove(c)
	if err != nil {
		fmt.Printf("%v", err)
		c.Abort()
		return
	}
	c.Next() // ok, pass to next handler
}

func (mw *XssMw) GetBlueMondayPolicy() *bluemonday.Policy {
	if mw.BmPolicy == "UGCPolicy" {
		return bluemonday.UGCPolicy()
	}

	return bluemonday.StrictPolicy()
}

func (mw *XssMw) XssRemove(c *gin.Context) error {
	ReqMethod := c.Request.Method
	ctHdr := c.Request.Header.Get("Content-Type")
	ctsLen := c.Request.Header.Get("Content-Length")
	ctLen, _ := strconv.Atoi(ctsLen)

	if ReqMethod == "POST" || ReqMethod == "PUT" {
		if ctLen > 1 && strings.Contains(ctHdr, "application/json") {
			err := mw.HandleJson(c)
			if err != nil {
				return err
			}
		} else if strings.Contains(ctHdr, "application/x-www-form-urlencoded") {
			err := mw.HandleXFormEncoded(c)
			if err != nil {
				return err
			}
		} else if strings.Contains(ctHdr, "multipart/form-data") {
			err := mw.HandleMultiPartFormData(c, ctHdr)
			if err != nil {
				return err
			}
		}
	} else if ReqMethod == "GET" {
		err := mw.HandleGETRequest(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mw *XssMw) HandleGETRequest(c *gin.Context) error {
	p := mw.GetBlueMondayPolicy()
	queryParams := c.Request.URL.Query()
	var fieldToSkip = map[string]bool{}
	for _, fts := range mw.FieldsToSkip {
		fieldToSkip[fts] = true
	}
	for key, items := range queryParams {
		if fieldToSkip[key] {
			continue
		}
		queryParams.Del(key)
		for _, item := range items {
			queryParams.Set(key, p.Sanitize(item))
		}
	}
	c.Request.URL.RawQuery = queryParams.Encode()
	return nil
}

func (mw *XssMw) HandleXFormEncoded(c *gin.Context) error {
	if c.Request.Body == nil {
		return nil
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(c.Request.Body); err != nil {
		return err
	}

	m, uerr := url.ParseQuery(buf.String())
	if uerr != nil {
		return uerr
	}

	p := mw.GetBlueMondayPolicy()

	var bq bytes.Buffer
	for k, v := range m {
		bq.WriteString(k)
		bq.WriteByte('=')

		var fndFld bool = false
		for _, fts := range mw.FieldsToSkip {
			if k == fts {
				bq.WriteString(url.QueryEscape(v[0]))
				fndFld = true
				break
			}
		}
		if !fndFld {
			bq.WriteString(url.QueryEscape(p.Sanitize(v[0])))
		}
		bq.WriteByte('&')
	}

	if bq.Len() > 1 {
		bq.Truncate(bq.Len() - 1) // remove last '&'
		bodOut := bq.String()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(bodOut)))
	} else {
		bufOut := buf.String()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(bufOut)))
	}

	return nil
}

func (mw *XssMw) HandleMultiPartFormData(c *gin.Context, ctHdr string) error {
	var ioreader io.Reader = c.Request.Body

	lenghtCtHdr := len(ctHdr)
	boundary := ctHdr[strings.Index(ctHdr, "boundary=")+9 : lenghtCtHdr]

	reader := multipart.NewReader(ioreader, boundary)

	var multiPrtFrm bytes.Buffer
	// unknown, so make up some param limit - 100 max should be enough
	for i := 0; i < 100; i++ {
		part, err := reader.NextPart()
		if err != nil {
			break
		}

		var buf bytes.Buffer
		n, err := io.Copy(&buf, part)
		if err != nil {
			return err
		}
		if n <= 0 {
			fmt.Println("xss.go.HandleMultiPartFormData -> error recreating Multipart form Request", n)
			// return errors.New("error recreating Multipart form Request")
		}
		multiPrtFrm.WriteString(`--` + boundary + "\r\n")
		if part.FileName() != "" {
			fn := part.FileName()
			mtype := part.Header.Get("Content-Type")
			multiPrtFrm.WriteString(`Content-Disposition: form-data; name="` + part.FormName() + "\"; ")
			multiPrtFrm.WriteString(`filename="` + fn + "\";\r\n")
			if mtype == "" {
				mtype = `application/octet-stream`
			}
			multiPrtFrm.WriteString(`Content-Type: ` + mtype + "\r\n\r\n")
			multiPrtFrm.WriteString(buf.String() + "\r\n")
		} else {
			multiPrtFrm.WriteString(`Content-Disposition: form-data; name="` + part.FormName() + "\";\r\n\r\n")
			p := bluemonday.StrictPolicy()
			if part.FormName() == "password" {
				multiPrtFrm.WriteString(buf.String() + "\r\n")
			} else {
				multiPrtFrm.WriteString(p.Sanitize(buf.String()) + "\r\n")
			}
		}
	}
	multiPrtFrm.WriteString("--" + boundary + "--\r\n")

	multiOut := multiPrtFrm.String()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(multiOut)))

	return nil
}

func (mw *XssMw) HandleJson(c *gin.Context) error {
	jsonBod, err := decodeJson(c.Request.Body)
	if err != nil {
		return err
	}

	buff, err := mw.jsonToStringMap(bytes.Buffer{}, jsonBod)
	if err != nil {
		return err
	}
	err = mw.SetRequestBodyJson(c, buff)
	if err != nil {
		return errors.New("set c.Request.Body Error")
	}
	return nil
}

func decodeJson(content io.Reader) (interface{}, error) {
	var jsonBod interface{}
	d := json.NewDecoder(content)
	d.UseNumber()
	err := d.Decode(&jsonBod)
	if err != nil {
		return nil, err
	}
	return jsonBod, err
}

func (mw *XssMw) jsonToStringMap(buff bytes.Buffer, jsonBod interface{}) (bytes.Buffer, error) {
	switch jbt := jsonBod.(type) {
	case map[string]interface{}:
		xmj := jsonBod.(map[string]interface{})
		var sbuff bytes.Buffer
		buff := mw.ConstructJson(xmj, sbuff)
		return buff, nil
	case []interface{}:
		var multiRec bytes.Buffer
		multiRec.WriteByte('[')
		for _, n := range jbt {
			xmj := n.(map[string]interface{})
			var sbuff bytes.Buffer
			buff = mw.ConstructJson(xmj, sbuff)
			multiRec.WriteString(buff.String())
			multiRec.WriteByte(',')
		}
		multiRec.Truncate(multiRec.Len() - 1) // remove last ','
		multiRec.WriteByte(']')
		return multiRec, nil
	}
	return bytes.Buffer{}, errors.New("unknown error")
}

func (mw *XssMw) SetRequestBodyJson(c *gin.Context, buff bytes.Buffer) error {
	bodOut := buff.String()

	enc := json.NewEncoder(ioutil.Discard)
	if merr := enc.Encode(&bodOut); merr != nil {
		return merr
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(bodOut)))
	return nil
}

func (mw *XssMw) ConstructJson(xmj XssMwJson, buff bytes.Buffer) bytes.Buffer {
	buff.WriteByte('{')

	p := mw.GetBlueMondayPolicy()

	m := xmj
	for k, v := range m {
		buff.WriteByte('"')
		buff.WriteString(k)
		buff.WriteByte('"')
		buff.WriteByte(':')

		// do fields to skip
		var fndFld bool = false
		for _, fts := range mw.FieldsToSkip {
			if string(k) == fts {
				buff.WriteString(fmt.Sprintf("%q", v))
				buff.WriteByte(',')
				fndFld = true
				break
			}
		}
		if fndFld {
			continue
		}

		var b bytes.Buffer
		apndBuff := mw.buildJsonApplyPolicy(v, b, p)
		buff.WriteString(apndBuff.String())
	}
	buff.Truncate(buff.Len() - 1) // remove last ','
	buff.WriteByte('}')

	return buff
}

func (mw *XssMw) buildJsonApplyPolicy(interf interface{}, buff bytes.Buffer, p *bluemonday.Policy) bytes.Buffer {
	switch v := interf.(type) {
	case map[string]interface{}:
		var sbuff bytes.Buffer
		scnd := mw.ConstructJson(v, sbuff)
		buff.WriteString(scnd.String())
		buff.WriteByte(',')
	case []interface{}:
		b := mw.unravelSlice(v, p)
		buff.WriteString(b.String())
		buff.WriteByte(',')
	case json.Number:
		buff.WriteString(p.Sanitize(fmt.Sprintf("%v", v)))
		buff.WriteByte(',')
	case string:
		buff.WriteString(fmt.Sprintf("%q", p.Sanitize(v)))
		buff.WriteByte(',')
	case float64:
		buff.WriteString(p.Sanitize(strconv.FormatFloat(v, 'g', 0, 64)))
		buff.WriteByte(',')
	default:
		if v == nil {
			buff.WriteString("null")
			buff.WriteByte(',')
		} else {
			buff.WriteString(p.Sanitize(fmt.Sprintf("%v", v)))
			buff.WriteByte(',')
		}
	}
	return buff
}

func (mw *XssMw) unravelSlice(slce []interface{}, p *bluemonday.Policy) bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte('[')
	for _, n := range slce {
		switch nn := n.(type) {
		case map[string]interface{}:
			var sbuff bytes.Buffer
			scnd := mw.ConstructJson(nn, sbuff)
			buff.WriteString(scnd.String())
			buff.WriteByte(',')
		case string:
			buff.WriteString(fmt.Sprintf("%q", p.Sanitize(nn)))
			buff.WriteByte(',')
		}
	}
	buff.Truncate(buff.Len() - 1) // remove last ','
	buff.WriteByte(']')
	return buff
}
