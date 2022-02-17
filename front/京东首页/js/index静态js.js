import { getToken } from "./module.js"

// 商品页面
;
(async function() {
    // console.log(JSON.parse(localStorage.getItem("Token")));
    let goods_msg_list = await getToken("/shop/all", localStorage.getItem("Token"))
    console.log(goods_msg_list);

})()
// 商品页面







// 判断用户状态
// console.log(JSON.parse(localStorage.getItem("Token")));
// if (JSON.parse(localStorage.getItem("Token"))) {
//     let login_register = document.querySelector(".login_register").querySelector("li")
//     let userForm = document.querySelector(".banner").querySelector(".text").querySelector(".userForm").querySelectorAll("a")

//     //访问他的孩子们
//     (async function() {
//         let user_form = await getToken(url, JSON.parse(localStorage.getItem("Token")))
//     })()
// }


//判断用户状态 


// second kill  部分
let x = 12
let time_list = document.querySelector(".seconds_kill").querySelectorAll("i")
let timer = setInterval(() => {
    let data = +new Date(`2022-2-${x} 22:00:00`)
    let now = +new Date()
    let h = Math.floor((data - now) / 1000 / 60 / 60 % 60)
    let m = Math.floor((data - now) / 1000 / 60 % 60)
    let s = Math.floor((data - now) / 1000 % 60);
    if (h <= 0 && s <= 0 && m <= 0) { x++ }
    time_list[4].innerText = s >= 10 ? s : "0" + s
    time_list[2].innerText = m >= 10 ? m : "0" + m
    time_list[0].innerText = h >= 10 ? h : "0" + h
}, 1);

// second kill   部分
//  特价销售 部分
let li_list = document.querySelector(".recommend").querySelector("ul").querySelectorAll("li")
let ul = document.querySelector(".recommend").querySelector("ul")
let goods_list = document.querySelector(".recommend").querySelectorAll(".goods")
let a_list = document.querySelector(".recommend").querySelector("ul").querySelectorAll("a")
    //委托给 ul办理这个 事件
a_list.forEach((e, index, arr) => {
    e.name = index

});
goods_list.forEach((e, index, arr) => {
    e.name = index

});

ul.addEventListener("mouseenter", (e) => {

        if (e.target.tagName.toLowerCase() != "a") return;
        for (let i = 0; i < goods_list.length; i++) {
            goods_list[i].style.display = "none"
            if (e.target.getAttribute("name") == goods_list[i].name)
                goods_list[i].style.display = ""
            console.log(goods_list[i].style.display);
        }

    }, true)
    // second kill   部分