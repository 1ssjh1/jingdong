import { getMessage } from "../module/module.js";

let btn = document.querySelector(".next")


btn.addEventListener("click", () => {
    console.log(111);
    let phone = document.querySelector(".message").children[1],
        phoneValue = phone.value;
    console.log(phoneValue);
    (async() => {
        let res = await getMessage(phoneValue)
        console.log(res);
        if (res.state == true) {
            window.location.href = "http://127.0.0.1:5500/code.html"
        }

    })()

})