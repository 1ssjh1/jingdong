import {get, post, postForm, getMessage, postFormDate } from "./mod.js"




let warning = document.querySelector(".warning") //获取 手机号码警告
let message = document.querySelector(".message").querySelector("input") //获取 手机号码

let codeWarning = document.querySelector(".code-warning") //获取验证码警告
let verifyCode = document.querySelector(".verify-code").querySelector("input") //获取 验证码窗口

let btn = document.querySelector(".next").children[0] // 获取进入下一个注册页面的节点
let verify = document.querySelector(".verify").children[0] // 获取验证请求按钮

// 聚焦input显现 提示信息
message.addEventListener("focus", () => {
    warning.style.opacity = "1"
})
verifyCode.addEventListener("focus", () => {
    codeWarning.style.opacity = "1"
})
console.log(verify);


// 验证手机号输入

verify.addEventListener("click", function() {
    let phoneValue = message.value;
    if (phoneValue == '') {
        warning.children[0].innerHTML = "&#xeb65; 啊你没有输手机号😓";
        warning.style.color = "red";
    }
    //判断手机号码是否正确
    if (phoneValue != "") {
        if (!(/^1[3456789]\d{9}$/.test(phoneValue))) {
            warning.children[0].innerHTML = "&#xeb65; 请填写正确的手机号码！😓(；′⌒`)";
            warning.style.color = "red";
        } else {
            localStorage.setItem('number', JSON.stringify(phoneValue))
            warning.children[0].innerHTML = "🐂🍺手机号码可以使用！😓o(*￣▽￣*)ブ";
            warning.style.color = "green";


            (async(phoneValue) => {
                let getCode = await getMessage(phoneValue)
                console.log(getCode);
            })()

            // let anser = JSON.parse(localStorage.getItem("number"))
            // console.log(anser);
            // localStorage.setItem('number', JSON.stringify(phoneValue))

        }
    }
})








// 验证 “验证码” 输入

btn.addEventListener("click", function() {
    let codeValue = verifyCode.value;
    if (codeValue == '') {

        codeWarning.children[0].innerHTML = "&#xeb65; 你没有输入验证码😓";

        codeWarning.style.color = "red";
    }
    if (codeValue != "") {
        if (!(/^\d{4}$/.test(codeValue))) {
            codeWarning.children[0].innerHTML = "&#xeb65; 请填写四位数字！😓(；′⌒`)";
            codeWarning.style.color = "red";
        } else {
            codeWarning.children[0].innerHTML = "🐂🍺芝麻开门";
            codeWarning.style.color = "green";
            // 跳转页面
            window.location.href = "./register-seceond.html"
        }
    }

})