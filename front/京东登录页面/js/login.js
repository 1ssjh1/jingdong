import {get, post, postForm, getMessage } from "./mod.js"
window.onload = function() {
    let username = document.querySelector(".username").querySelector("input")

    let code = document.querySelector(".code").querySelector("input")

    let enterBtn = document.querySelector(".enter-btn").querySelector("button")


    enterBtn.addEventListener("click", (e) => {
        e.preventDefault()
        let usernameValue = username.value
        let codeValue = code.value
        let formData = {
                username: usernameValue,
                password: codeValue

            }
            // sianao 
            // 1000
        ;
        (async() => {
            let res = await postForm("/login", formData)
            console.log(res);
            if (res.state == true) {
                console.log(res);
                console.log(res.Token);
                localStorage.setItem('Token', res.Token)
                    // http://127.0.0.1:5502/html/index.html
                    //http://127.0.0.1:5502/html/index.html   首页
                    // http://127.0.0.1:5502/html/index.html
                    // window.location.href = "http://127.0.0.1:5500/index.html"
            }

        })()

    })







}