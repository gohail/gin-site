
$(document).ready(function () {





    function headerAnimateText() {
        setTimeout('$(".animateMainText0").animate({ opacity: 1, }, 1000 )', 000)
        setTimeout('$(".animateMainText1").animate({ opacity: 1, }, 1000 )', 500)
        setTimeout('$(".animateMainText2").animate({ opacity: 1, }, 1000 )', 1000)
        setTimeout('$(".animateMainText3").animate({ opacity: 1, }, 1000 )', 1500)
    }
    headerAnimateText();
    /*
    $(window).on("scroll", function () {
        var scrollTopS;
        scrollTopS = $(window).scrollTop();
        if (scrollTopS > 200) {
            headerAnimateTextHide();
        } else if (scrollTopS < 20) {
            headerAnimateTextHide();
        }
    });
    */
});
/**********************************************************************************/

/*скрипт для главной страницы блок наши работы меняет массив с фотками*/
/**********************************************************************************/
/*
$(document).ready(function () {
    var imgHead = [
        "user_guide/_images/ourWorks1.jpg",
        "user_guide/_images/ourWorks2.jpg"
    ], i = 1;

    function csaHead() {
        if (i > (imgHead.length - 1)) {
            i = 1;
            $('#ourWorksMainPage').css({
                'background': 'url(' + imgHead[0] + ')',
                'background-size': 'cover',
                'background-position': 'center'
            });
        } else {
            $('#ourWorksMainPage').css({
                'background': 'url(' + imgHead[i] + ')',
                'background-size': 'cover',
                'background-position': 'center'
            });
            i++;
        }
    }
    var intervalCsaHead = setInterval(csaHead, 4000);
});
*/
/********************************************************************************************/

/*модальное окно собирается в зависимости от атрибутов кнопки*/
/*********************************************************************************************/
$(function () {
    /* Программно создаём модальные окна */
    var showModal1 = false;
    // создаём объект "Модальное окно 1"
    var myModal1 = new ModalApp.ModalProcess({
        id: 'myModal1',
        title: 'Что Вас интересует?',
        body: ''
    });
    myModal1.init();
    /* Определяем поведение окон */
    // при нажатии на кнопку с классом modal-call-back
    $('.modal-call-back').click(function () {
        /*как рабоатет скрипт
        в атрибутах кнопки есть 3 разных параметра. мейл, пхоне, и месседж. в фоме будут отобразаться только те
        которые указаны. если параметр месседж не указан значит в форме не будет формы для ввода сообщения.
        */

        var contentBody = '';
        var contentFooter = '';
        if($(this).attr('data-title') ){ var title = $(this).attr('data-title'); }

        if($(this).attr('data-phone')){
            contentBody += "<div class=\"form-group\"> " +
                "<label for=\"inputName\"  class=\"control-label\">Телефон</label>\n" +
                "<input type=\"text\"  class=\"form-control\" id=\"\" placeholder=\"Введите номер\" data-error=\"Введите номер телефона\" required>\n" +
                "    <div class=\"help-block with-errors\"></div>\n" +
                "</div>";
            }
        if($(this).attr('data-mail')){
            contentBody += "<div class=\"form-group\">\n" +
                "    <label for=\"inputEmail\" class=\"control-label\">Email</label>\n" +
                "    <input type=\"email\" class=\"form-control\" id=\"inputEmail\" placeholder=\"mail@host.dom\" data-error=\"Введите почтовый адрес\" required>\n" +
                "    <div class=\"help-block with-errors\"></div>\n" +
                "  </div>";
        }
        if($(this).attr('data-mess')){
            contentBody += "<div class=\"form-group\">\n" +
                "    <label for=\"inputText\" class=\"control-label\">Повідомлення</label>\n" +
                "    <textarea class=\"form-control\" rows=\"3\" id=\"inputMessage\" placeholder=\"Ваше Сообщение\"></textarea>\n" +
                "    <div class=\"help-block with-errors\"></div>\n" +
                "  </div>";


        }
        contentBody = "<form  id=\"myform111\"  >"+contentBody+'<div class=\"form-group\"> <button  id=\"btn_sand_masusage\" class="btn btn-success">Отправить</button> </div></form>';
        contentFooter = '';
        myModal1.changeTitle(title);
        myModal1.changeBody(contentBody);
        myModal1.changeFooter(contentFooter);

        myModal1.showModal(); // сздаем окно с формой
        $('#myModal1').on('shown.bs.modal', function () { $('#inputMessage').focus(); });
        $(function () { $(".inputPhone").mask("+38(999) 999-9999");  });//Маска для ввода телефона
        $('#myform111').validator();//запускаем валидацию

        /**********************************************************************************************/

        $('form').submit(function (e) {
            //проверяем есить ли ошибки в форме после валидации если нет то пытаемся отправить
            if (!e.isDefaultPrevented()) {
                e.preventDefault(); // отмена действия сабмита по умолчанию


                var pageTitle = document.title; // страница с которой прийдет запрос
                var pushButton = title; // кнопка с которой прийдет запрос
                var base_url = location.origin; // Текущий домен


                if ($("#inputPhone").length){var userPhone = document.getElementById("inputPhone").value;}
                if ($("#inputEmail").length){var userEmail = document.getElementById("inputEmail").value;}
                if ($("#inputMessage").length){var userMessage = document.getElementById("inputMessage").value;}

                $.ajax({
                    type: "post",
                    url: base_url+'/mail/getphone',
                    //dataType: 'json',
                    data: {
                        userPage: pageTitle,
                        pushButton: pushButton,
                        userPhone: userPhone,
                        userEmail: userEmail,
                        userMessage: userMessage
                    },
                    beforeSend: function(){
                        myModal1.changeBody("<br><div class='text-center'>Отправка... </div><br>");
                    },
                    success: function (data) {
                        myModal1.changeBody("<br>" +data+ "<br>");
                        //setTimeout(function() { setTimeout("$(\"#myModal1\").modal('hide')", 1000); });
                    },
                });
            }
        });
    });
});
/*конструктор модального окна закончен*/
/*********************************************************************************************/


/*Скрипт прокрутки страницы по якорю начался*/
/***********************************************************************************************/

$(document).ready(function() {
  $("a.skrollto").click(function() {
    var elementClick = $(this).attr("href")
    var destination = $(elementClick).offset().top;
    jQuery("html:not(:animated),body:not(:animated)").animate({
      scrollTop: destination
    }, 800);

    return false;
  });

  $("#go_to_order").click(function() {
    $('#name').focus();
  });

});

/*Скрипт прокрутки страницы по якорю хзакончен*/
/***********************************************************************************************/

$( document ).ready(function() {
    $('.btnModalCart').on('click', function() {
        $("#exampleModal").modal();
    });
});


/*Конструктор модального окна --/ */
var ModalApp = {};
ModalApp.ModalProcess = function (parameters) {
    this.id = parameters['id'] || 'modal';
    this.selector = parameters['selector'] || '';
    this.title = parameters['title'] || 'Заголовок модального окна';
    this.body = parameters['body'] || 'Содержимое модального окна';
    this.footer = parameters['footer'] || '<button type="button" class="btn btn-default" data-dismiss="modal">Закрыть</button>';
    this.content = '<div id="'+this.id+'" class="modal fade"  tabindex="-1" role="dialog">'+
        '<div class="modal-dialog" role="document">'+
        '<div class="modal-content" >'+

        //'<form>'+
        '<div class="modal-header" style="background: rgb(140, 62, 172); font-weight: 500; color: white;">'+
        '<h4 class="modal-title">'+this.title+'</h4>'+
        '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>'+

        '</div>'+
        '<div class="modal-body" style="padding: 15px; color: gray;"> <form id="myform111" data-toggle="validator" role="form">'+this.body+' </form> </div>'+
        //'<div class="modal-footer">'+this.footer+'</div>'+
        //'</form>'+
        '</div>'+
        '</div>'+
        '</div>';
    //this.content = '<form id="myform111" data-toggle="validator" role="form">'+this.content+'</form>';

    this.init = function() {
        if ($('#'+this.id).length==0) {
            $('body').prepend(this.content);
        }
        if (this.selector) {
            $(document).on('click',this.selector, $.proxy(this.showModal,this));
        }
    }
}
ModalApp.ModalProcess.prototype.changeTitle = function(content) {
    $('#' + this.id + ' .modal-title').html(content);
};
ModalApp.ModalProcess.prototype.changeBody = function(content) {
    $('#' + this.id + ' .modal-body' + '').html(content);
};
ModalApp.ModalProcess.prototype.changeFooter = function(content) {
    $('#' + this.id + ' .modal-footer').html(content);
};
ModalApp.ModalProcess.prototype.showModal = function() {
    $('#' + this.id).modal('show');
};
ModalApp.ModalProcess.prototype.hideModal = function() {
    $('#' + this.id).modal('hide');
};
ModalApp.ModalProcess.prototype.updateModal = function() {
    $('#' + this.id).modal('handleUpdate');
};
/* -- /Конструктор модального окна*/

/*Маска телефона -- / */
/*
    jQuery Masked Input Plugin
    Copyright (c) 2007 - 2015 Josh Bush (digitalbush.com)
    Licensed under the MIT license (http://digitalbush.com/projects/masked-input-plugin/#license)
    Version: 1.4.1
*/
!function(factory) {
    "function" == typeof define && define.amd ? define([ "jquery" ], factory) : factory("object" == typeof exports ? require("jquery") : jQuery);
}(function($) {
    var caretTimeoutId, ua = navigator.userAgent, iPhone = /iphone/i.test(ua), chrome = /chrome/i.test(ua), android = /android/i.test(ua);
    $.mask = {
        definitions: {
            "9": "[0-9]",
            a: "[A-Za-z]",
            "*": "[A-Za-z0-9]"
        },
        autoclear: !0,
        dataName: "rawMaskFn",
        placeholder: "_"
    }, $.fn.extend({
        caret: function(begin, end) {
            var range;
            if (0 !== this.length && !this.is(":hidden")) return "number" == typeof begin ? (end = "number" == typeof end ? end : begin,
                this.each(function() {
                    this.setSelectionRange ? this.setSelectionRange(begin, end) : this.createTextRange && (range = this.createTextRange(),
                        range.collapse(!0), range.moveEnd("character", end), range.moveStart("character", begin),
                        range.select());
                })) : (this[0].setSelectionRange ? (begin = this[0].selectionStart, end = this[0].selectionEnd) : document.selection && document.selection.createRange && (range = document.selection.createRange(),
                begin = 0 - range.duplicate().moveStart("character", -1e5), end = begin + range.text.length),
                {
                    begin: begin,
                    end: end
                });
        },
        unmask: function() {
            return this.trigger("unmask");
        },
        mask: function(mask, settings) {
            var input, defs, tests, partialPosition, firstNonMaskPos, lastRequiredNonMaskPos, len, oldVal;
            if (!mask && this.length > 0) {
                input = $(this[0]);
                var fn = input.data($.mask.dataName);
                return fn ? fn() : void 0;
            }
            return settings = $.extend({
                autoclear: $.mask.autoclear,
                placeholder: $.mask.placeholder,
                completed: null
            }, settings), defs = $.mask.definitions, tests = [], partialPosition = len = mask.length,
                firstNonMaskPos = null, $.each(mask.split(""), function(i, c) {
                "?" == c ? (len--, partialPosition = i) : defs[c] ? (tests.push(new RegExp(defs[c])),
                null === firstNonMaskPos && (firstNonMaskPos = tests.length - 1), partialPosition > i && (lastRequiredNonMaskPos = tests.length - 1)) : tests.push(null);
            }), this.trigger("unmask").each(function() {
                function tryFireCompleted() {
                    if (settings.completed) {
                        for (var i = firstNonMaskPos; lastRequiredNonMaskPos >= i; i++) if (tests[i] && buffer[i] === getPlaceholder(i)) return;
                        settings.completed.call(input);
                    }
                }
                function getPlaceholder(i) {
                    return settings.placeholder.charAt(i < settings.placeholder.length ? i : 0);
                }
                function seekNext(pos) {
                    for (;++pos < len && !tests[pos]; ) ;
                    return pos;
                }
                function seekPrev(pos) {
                    for (;--pos >= 0 && !tests[pos]; ) ;
                    return pos;
                }
                function shiftL(begin, end) {
                    var i, j;
                    if (!(0 > begin)) {
                        for (i = begin, j = seekNext(end); len > i; i++) if (tests[i]) {
                            if (!(len > j && tests[i].test(buffer[j]))) break;
                            buffer[i] = buffer[j], buffer[j] = getPlaceholder(j), j = seekNext(j);
                        }
                        writeBuffer(), input.caret(Math.max(firstNonMaskPos, begin));
                    }
                }
                function shiftR(pos) {
                    var i, c, j, t;
                    for (i = pos, c = getPlaceholder(pos); len > i; i++) if (tests[i]) {
                        if (j = seekNext(i), t = buffer[i], buffer[i] = c, !(len > j && tests[j].test(t))) break;
                        c = t;
                    }
                }
                function androidInputEvent() {
                    var curVal = input.val(), pos = input.caret();
                    if (oldVal && oldVal.length && oldVal.length > curVal.length) {
                        for (checkVal(!0); pos.begin > 0 && !tests[pos.begin - 1]; ) pos.begin--;
                        if (0 === pos.begin) for (;pos.begin < firstNonMaskPos && !tests[pos.begin]; ) pos.begin++;
                        input.caret(pos.begin, pos.begin);
                    } else {
                        for (checkVal(!0); pos.begin < len && !tests[pos.begin]; ) pos.begin++;
                        input.caret(pos.begin, pos.begin);
                    }
                    tryFireCompleted();
                }
                function blurEvent() {
                    checkVal(), input.val() != focusText && input.change();
                }
                function keydownEvent(e) {
                    if (!input.prop("readonly")) {
                        var pos, begin, end, k = e.which || e.keyCode;
                        oldVal = input.val(), 8 === k || 46 === k || iPhone && 127 === k ? (pos = input.caret(),
                            begin = pos.begin, end = pos.end, end - begin === 0 && (begin = 46 !== k ? seekPrev(begin) : end = seekNext(begin - 1),
                            end = 46 === k ? seekNext(end) : end), clearBuffer(begin, end), shiftL(begin, end - 1),
                            e.preventDefault()) : 13 === k ? blurEvent.call(this, e) : 27 === k && (input.val(focusText),
                            input.caret(0, checkVal()), e.preventDefault());
                    }
                }
                function keypressEvent(e) {
                    if (!input.prop("readonly")) {
                        var p, c, next, k = e.which || e.keyCode, pos = input.caret();
                        if (!(e.ctrlKey || e.altKey || e.metaKey || 32 > k) && k && 13 !== k) {
                            if (pos.end - pos.begin !== 0 && (clearBuffer(pos.begin, pos.end), shiftL(pos.begin, pos.end - 1)),
                                    p = seekNext(pos.begin - 1), len > p && (c = String.fromCharCode(k), tests[p].test(c))) {
                                if (shiftR(p), buffer[p] = c, writeBuffer(), next = seekNext(p), android) {
                                    var proxy = function() {
                                        $.proxy($.fn.caret, input, next)();
                                    };
                                    setTimeout(proxy, 0);
                                } else input.caret(next);
                                pos.begin <= lastRequiredNonMaskPos && tryFireCompleted();
                            }
                            e.preventDefault();
                        }
                    }
                }
                function clearBuffer(start, end) {
                    var i;
                    for (i = start; end > i && len > i; i++) tests[i] && (buffer[i] = getPlaceholder(i));
                }
                function writeBuffer() {
                    input.val(buffer.join(""));
                }
                function checkVal(allow) {
                    var i, c, pos, test = input.val(), lastMatch = -1;
                    for (i = 0, pos = 0; len > i; i++) if (tests[i]) {
                        for (buffer[i] = getPlaceholder(i); pos++ < test.length; ) if (c = test.charAt(pos - 1),
                                tests[i].test(c)) {
                            buffer[i] = c, lastMatch = i;
                            break;
                        }
                        if (pos > test.length) {
                            clearBuffer(i + 1, len);
                            break;
                        }
                    } else buffer[i] === test.charAt(pos) && pos++, partialPosition > i && (lastMatch = i);
                    return allow ? writeBuffer() : partialPosition > lastMatch + 1 ? settings.autoclear || buffer.join("") === defaultBuffer ? (input.val() && input.val(""),
                        clearBuffer(0, len)) : writeBuffer() : (writeBuffer(), input.val(input.val().substring(0, lastMatch + 1))),
                        partialPosition ? i : firstNonMaskPos;
                }
                var input = $(this), buffer = $.map(mask.split(""), function(c, i) {
                    return "?" != c ? defs[c] ? getPlaceholder(i) : c : void 0;
                }), defaultBuffer = buffer.join(""), focusText = input.val();
                input.data($.mask.dataName, function() {
                    return $.map(buffer, function(c, i) {
                        return tests[i] && c != getPlaceholder(i) ? c : null;
                    }).join("");
                }), input.one("unmask", function() {
                    input.off(".mask").removeData($.mask.dataName);
                }).on("focus.mask", function() {
                    if (!input.prop("readonly")) {
                        clearTimeout(caretTimeoutId);
                        var pos;
                        focusText = input.val(), pos = checkVal(), caretTimeoutId = setTimeout(function() {
                            input.get(0) === document.activeElement && (writeBuffer(), pos == mask.replace("?", "").length ? input.caret(0, pos) : input.caret(pos));
                        }, 10);
                    }
                }).on("blur.mask", blurEvent).on("keydown.mask", keydownEvent).on("keypress.mask", keypressEvent).on("input.mask paste.mask", function() {
                    input.prop("readonly") || setTimeout(function() {
                        var pos = checkVal(!0);
                        input.caret(pos), tryFireCompleted();
                    }, 0);
                }), chrome && android && input.off("input.mask").on("input.mask", androidInputEvent),
                    checkVal();
            });
        }
    });
});
/* /--Маска телефона */

/*счетчик акция */
function getTimeRemaining(endtime) {
    var t = Date.parse(endtime) - Date.parse(new Date());
    var seconds = Math.floor((t / 1000) % 60);
    var minutes = Math.floor((t / 1000 / 60) % 60);
    var hours = Math.floor((t / (1000 * 60 * 60)) % 24);
    var days = Math.floor(t / (1000 * 60 * 60 * 24));
    return {
        'total': t,
        'days': days,
        'hours': hours,
        'minutes': minutes,
        'seconds': seconds
    };
}

function initializeClock(id, endtime) {
    var clock = document.getElementById(id);
    var daysSpan = clock.querySelector('.days');
    var hoursSpan = clock.querySelector('.hours');
    var minutesSpan = clock.querySelector('.minutes');
    var secondsSpan = clock.querySelector('.seconds');

    function updateClock() {
        var t = getTimeRemaining(endtime);

        daysSpan.innerHTML = t.days;
        hoursSpan.innerHTML = ('0' + t.hours).slice(-2);
        minutesSpan.innerHTML = ('0' + t.minutes).slice(-2);
        secondsSpan.innerHTML = ('0' + t.seconds).slice(-2);

        if (t.total <= 0) {
            clearInterval(timeinterval);
        }
    }

    updateClock();
    var timeinterval = setInterval(updateClock, 1000);
}

var deadline="January 01 2018 00:00:00 GMT+0300"; //for Ukraine
var deadline = new Date(Date.parse(new Date()) + 8 * 24 * 60 * 60 * 1000); // for endless timer
initializeClock('clockdiv', deadline);
/* /- счетчик акция */
