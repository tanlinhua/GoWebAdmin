// 顶部鼠标移上显示
$(".main_tips").mouseenter(function () {
    var str = $(this).find("a").attr("tips");
    layer.tips(str, $(this), {
        tips: [1, "#FF5722"],
        time: 1000,
    });
});

/**
 * 系统公告
 */
$(document).on("click", "#notice", noticeFun);
var noticeContent = "Hello World.";
getNoticeMsg();
function getNoticeMsg() {
    $.ajax({
        url: "/admin/message/get",
        type: "GET",
        dataType: "json",
        success: function (result) {
            if (result.code == 0) {
                noticeContent = result.data.msg;
            } else {
                layer.msg(result.msg);
            }
        },
        error: function () {
            layer.msg("获取message失败");
        },
    });
}
function noticeFun() {
    getNoticeMsg();
    layer.open({
        type: 0,
        title: "系统公告",
        btn: "我知道啦",
        btnAlign: "c",
        content: noticeContent,
        yes: function (index) {
            sessionStorage.setItem("notice", "true");
            layer.close(index);
        },
        cancel: function (index) {},
    });
}
!(function () {
    var notice = sessionStorage.getItem("notice"); // 是否已经提示过
    if (notice != "true") {
        setTimeout(function () {
            noticeFun();
        }, 3000);
    }
})();

/**
 * 全屏/退出全屏
 */
$("body").on("click", "#fullScreen", function () {
    if ($(this).children("i").hasClass("layui-icon-screen-restore")) {
        screenFun(2).then(function () {
            $("#fullScreen").children("i").eq(0).removeClass("layui-icon-screen-restore");
        });
    } else {
        screenFun(1).then(function () {
            $("#fullScreen").children("i").eq(0).addClass("layui-icon-screen-restore");
        });
    }
});

/**
 * 全屏和退出全屏的方法
 * @param num 1代表全屏 2代表退出全屏
 * @returns {Promise}
 */
function screenFun(num) {
    num = num || 1;
    num = num * 1;
    var docElm = document.documentElement;

    switch (num) {
        case 1:
            if (docElm.requestFullscreen) {
                docElm.requestFullscreen();
            } else if (docElm.mozRequestFullScreen) {
                docElm.mozRequestFullScreen();
            } else if (docElm.webkitRequestFullScreen) {
                docElm.webkitRequestFullScreen();
            } else if (docElm.msRequestFullscreen) {
                docElm.msRequestFullscreen();
            }
            break;
        case 2:
            if (document.exitFullscreen) {
                document.exitFullscreen();
            } else if (document.mozCancelFullScreen) {
                document.mozCancelFullScreen();
            } else if (document.webkitCancelFullScreen) {
                document.webkitCancelFullScreen();
            } else if (document.msExitFullscreen) {
                document.msExitFullscreen();
            }
            break;
    }

    return new Promise(function (res, rej) {
        res("返回值");
    });
}

var okUtils = {
    // localStorage 二次封装
    local: function (name, value) {
        if (value) {
            /**设置*/
            if (typeof value == "object") {
                localStorage.setItem(name, JSON.stringify(value));
            } else {
                localStorage.setItem(name, value);
            }
        } else if (null !== value) {
            /**获取*/
            let val = localStorage.getItem(name);
            try {
                val = JSON.parse(val);
                return val;
            } catch (err) {
                return val;
            }
        } else {
            /**清除*/
            return localStorage.removeItem(name);
        }
    },

    /**
     * 格式化当前日期
     * @param date
     * @param fmt
     * @returns {void | string}
     */
    dateFormat: function (date, fmt) {
        date = date || new Date();
        fmt = fmt || "yyyy年M月s日";
        var o = {
            "M+": date.getMonth() + 1,
            "d+": date.getDate(),
            "h+": date.getHours(),
            "m+": date.getMinutes(),
            "s+": date.getSeconds(),
            "q+": Math.floor((date.getMonth() + 3) / 3),
            S: date.getMilliseconds(),
        };
        if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
        for (var k in o) if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" + o[k]).length));
        return fmt;
    },

    // 字符串加密
    toCode: function (str) {
        //加密字符串
        //定义密钥，36个字母和数字
        var key = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        var len = key.length; //获取密钥的长度
        var a = key.split(""); //把密钥字符串转换为字符数组
        var s = "",
            b,
            b1,
            b2,
            b3; //定义临时变量
        for (var i = 0; i < str.length; i++) {
            //遍历字符串
            b = str.charCodeAt(i); //逐个提取每个字符，并获取Unicode编码值
            b1 = b % len; //求Unicode编码值得余数
            b = (b - b1) / len; //求最大倍数
            b2 = b % len; //求最大倍数的于是
            b = (b - b2) / len; //求最大倍数
            b3 = b % len; //求最大倍数的余数
            s += a[b3] + a[b2] + a[b1]; //根据余数值映射到密钥中对应下标位置的字符
        }
        return s; //返回这些映射的字符
    },
    // 字符串解密
    fromCode: function (str) {
        //定义密钥，36个字母和数字
        var key = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        var len = key.length; //获取密钥的长度
        var b,
            b1,
            b2,
            b3,
            d = 0,
            s; //定义临时变量
        s = new Array(Math.floor(str.length / 3)); //计算加密字符串包含的字符数，并定义数组
        b = s.length; //获取数组的长度
        for (var i = 0; i < b; i++) {
            //以数组的长度循环次数，遍历加密字符串
            b1 = key.indexOf(str.charAt(d)); //截取周期内第一个字符串，计算在密钥中的下标值
            d++;
            b2 = key.indexOf(str.charAt(d)); //截取周期内第二个字符串，计算在密钥中的下标值
            d++;
            b3 = key.indexOf(str.charAt(d)); //截取周期内第三个字符串，计算在密钥中的下标值
            d++;
            s[i] = b1 * len * len + b2 * len + b3; //利用下标值，反推被加密字符的Unicode编码值
        }
        b = eval("String.fromCharCode(" + s.join(",") + ")"); // 用fromCharCode()算出字符串
        return b; //返回被解密的字符串
    },
};

/**
 * 锁定账户
 */
var lock_inter = "";
showLockView(okUtils);
$("#lock").click(function () {
    let store = okUtils.local("local_password");
    if (!store) {
        layui.layer.open({
            type: 1,
            title: "请输入解锁密码",
            area: ["300px", "170px"],
            content: $("#lock_screen_view"),
            btn: ["确认", "取消"],
            yes: function (index, layero) {
                var input = $("#lock_input_password").val();
                if (!input) {
                    layui.layer.msg("请输入");
                    return;
                }
                okUtils.local("local_password", okUtils.toCode(input));
                okUtils.local("isLock", "1"); // 设置锁屏缓存防止刷新失效
                showLockView(okUtils); // 锁屏
                layer.close(index);
            },
        });
    } else {
        layui.layer.confirm("确定要锁定账户吗？", function (index) {
            layer.close(index);
            okUtils.local("isLock", "1"); // 设置锁屏缓存防止刷新失效
            showLockView(okUtils); // 锁屏
        });
    }
});

/**锁屏方法*/
function showLockView(okUtils) {
    let localLock = okUtils.local("isLock");
    $("#lockPassword").val("");
    if (!localLock) {
        return;
    }

    $(".lock-screen").show();
    Snowflake("snowflake"); // 雪花

    var lock_bgs = $(".lock-screen .lock-bg img");
    $(".lock-content .time .hhmmss").html(okUtils.dateFormat("", "hh <p lock='lock'>:</p> mm"));
    $(".lock-content .time .yyyymmdd").html(okUtils.dateFormat("", "yyyy 年 M 月 dd 日"));

    var i = 0,
        k = 0;
    lock_inter = setInterval(function () {
        i++;
        if (i % 8 == 0) {
            k = k + 1 >= lock_bgs.length ? 0 : k + 1;
            i = 0;
            lock_bgs.removeClass("active");
            $(lock_bgs[k]).addClass("active");
        }
        $(".lock-content .time .hhmmss").html(okUtils.dateFormat("", "hh <p lock='lock'>:</p> mm"));
    }, 1000);

    //提交密码
    layui.form.on("submit(lockSubmit)", function (data) {
        let store = okUtils.local("local_password");
        if (okUtils.toCode(data.field.lock_password) != store) {
            layer.msg("密码错误", {
                icon: 5,
                zIndex: 999999991,
            });
        } else {
            layer.msg("密码正确", {
                time: 1000,
                icon: 6,
                zIndex: 999999992,
                end: function () {
                    okUtils.local("isLock", null); //清除锁屏的缓存
                    $("#lockPassword").val(""); //清除输入框的密码
                    $(".lock-screen").hide();
                    clearInterval(lock_inter);
                },
            });
        }
        return false;
    });

    //退出登录
    $("#lockQuit").click(function () {
        okUtils.local("isLock", null); //清除锁屏的缓存
        okUtils.local("local_password", null);
        window.location.replace("logout?layui=1"); //替换当前页面
    });
}
