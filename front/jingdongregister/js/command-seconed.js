import {get, post, postForm, getMessage, postFormDate } from "./mod.js"






window.onload = function() {
    let username = document.querySelector(".username").querySelector("input") //用户名
    let username_warning = document.querySelector(".username-warning")

    let code = document.querySelector(".code").querySelector("input") //密码
    let code_warning = document.querySelector(".code-warning")

    let makeSureCode = document.querySelector(".makeSure-code").querySelector("input") //确认密码
    let makeSure_code_warning = document.querySelector(".makeSure-code-warning")

    let registerBtn = document.querySelector(".register-btn").children[0] //下一步按钮


    let verifyCode = document.querySelector(".verify-code").querySelector("input") //验证码
    let verify_code_warning = document.querySelector(".verify-code-warning")

    let input_list = document.querySelectorAll("input")




    let verifyCodeValue = verifyCode.value;







    input_list[0].addEventListener("blur", function(e) {
        let usernameValue = username.value;

        if (usernameValue == "") {
            username_warning.children[0].innerHTML = "&#xeb65 用户名不能为空"
            username_warning.children[0].style.color = "red"

        } else if (!(/^[a-zA-Z0-9_-]{4,8}$/).test(usernameValue)) {
            username_warning.children[0].innerHTML = "&#xeb65;用户名不符规范！（大小写字母3—8位）"
            username_warning.children[0].style.color = "red"

        } else {
            username_warning.children[0].innerHTML = "&#xeb65; G O O D!"
            username_warning.children[0].style.color = "green"
        }
    })


    input_list[2].addEventListener("blur", function(e) {

        let codeValue = code.value;

        if (codeValue == "") {
            code_warning.children[0].innerHTML = "&#xeb65 密码不能为空(⊙o⊙)？"
            code_warning.children[0].style.color = "red"
            return
        } else if (!(/^[a-zA-Z0-9_-]{4,8}$/).test(codeValue)) {
            code_warning.children[0].innerHTML = "&#xeb65 密码不符合规范(╯▔皿▔)╯"
            code_warning.children[0].style.color = "red"
            return
        } else {
            code_warning.children[0].innerHTML = "&#xeb65 G O O D"
            code_warning.children[0].style.color = "green"
        }
    })

    input_list[3].addEventListener("blur", function(e) {
        let makeSureCodeValue = makeSureCode.value
        let codeValue = code.value;
        if (codeValue != makeSureCodeValue) {
            makeSure_code_warning.children[0].innerHTML = "&#xeb65 前后密码不一致哦"
            makeSure_code_warning.children[0].style.color = "red"
            return

        } else {
            makeSure_code_warning.children[0].innerHTML = "&#xeb65; G O O D!"
            makeSure_code_warning.children[0].style.color = "green"
            return
        }
    })



    registerBtn.addEventListener("click", () => {
        if (makeSure_code_warning.children[0].style.color == "green" && code_warning.children[0].style.color == "green" && username_warning.children[0].style.color == "green") {
            let makeSureCodeValue = makeSureCode.value,
                codeValue = code.value,
                usernameValue = username.value,
                verifyCodeValue = verifyCode.value,
                number = JSON.parse(localStorage.getItem("number"));

            let data = {
                username: usernameValue,
                password: codeValue,
                number: number,
                code: verifyCodeValue
            }

            ;
            (async() => {
                let res = await postForm("/register", data)
                console.log(res);
                console.log(res.state);

                if (res.state == false) {
                    verify_code_warning.children[0].innerHTML = `${res.msg}`
                    verify_code_warning.children[0].style.color = "red"
                } else if (res.state == true) {
                    localStorage.setItem("username", JSON.stringify(usernameValue))
                    console.log(localStorage.getItem("username"));

                    window.location.href = "./register-succeed.html"
                }
            })()


        }




    })



    // // formDate.set("username", JSON.stringify(usernameValue));
    // // formDate.set("password", JSON.stringify(codeValue));
    // // formDate.set("code", JSON.stringify(verifyCodeValue));
    // // formDate.set("number", JSON.stringify(number));


















}