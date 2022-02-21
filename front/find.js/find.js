import { find } from "../module/module.js";


let register_btn = document.querySelector(".register-btn")


register_btn.addEventListener("click", () => {


    let user = document.querySelector(".username").children[1],
        verify_code = document.querySelector(".verify-code").children[1],
        phone = document.querySelector(".code").children[1],
        code = document.querySelector(".makeSure-code").children[1]
    console.log(user, verify_code, phone, code);

    let formData = new FormData()
    formData.append("username", user.value)
    formData.append("password", code.value)
    formData.append("code", verify_code.value)
    formData.append("number", phone.value)

    ;
    (async() => {
        let res = await find("find", formData)
        console.log(res);

    })()


})