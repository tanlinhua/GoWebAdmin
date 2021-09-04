layui.define(['element'], function (exports) {

    var $ = layui.$;
    $('.input-field').on('change', function () {
        var $this = $(this);
        var value = $.trim($this.val());
        var $parent = $this.parent();

        if (value !== '') {
            $parent.addClass('field-focus');
        } else {
            $parent.removeClass('field-focus');
        }
    });
    exports('login');
});