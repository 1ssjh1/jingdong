import {get, post, postForm, getMessage, postFormDate } from "./mod.js"






window.onload = function() {
    let username = document.querySelector(".username").querySelector("input") //用户名


    let code = document.querySelector(".code").querySelector("input") //密码


    let makeSureCode = document.querySelector(".makeSure-code").querySelector("input") //确认密码


    let registerBtn = document.querySelector(".register-btn").children[0] //下一步按钮


    let verifyCode = document.querySelector(".verify-code").querySelector("input") //验证码



    registerBtn.addEventListener("click", function() {
        let usernameValue = username.value;
        let codeValue = code.value;
        let verifyCodeValue = verifyCode.value;
        let number = JSON.parse(localStorage.getItem("number"));
        let data = {
            username: usernameValue,
            password: codeValue,
            number: number,
            code: verifyCodeValue
        }

        // formDate.set("username", JSON.stringify(usernameValue));
        // formDate.set("password", JSON.stringify(codeValue));
        // formDate.set("code", JSON.stringify(verifyCodeValue));
        // formDate.set("number", JSON.stringify(number));

        ;
        (async() => {
            let res = await postForm("/register", data)
            console.log(res);
            console.log(res.state);
            if (res.state == false) {


                localStorage.setItem("username", JSON.stringify(usernameValue))
                console.log(localStorage.getItem("username"));

                window.location.href = "./register-succeed.html"
            }
        })()






    })











}