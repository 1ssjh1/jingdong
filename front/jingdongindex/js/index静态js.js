import { getToken, out, get, get_news } from "./module.js"
//判断用户状态
console.log(localStorage.getItem("Token"));

// 删除 商品
let listening_code = false


function listening() {
    for (let i = 0; i < 15; i++) {
        if (!listening_code) {
            localStorage.removeItem(`Gid${i}_src`)
            localStorage.removeItem(`Gid${i}_price`)
            localStorage.removeItem(`Gid${i}_id`)
            localStorage.removeItem(`Gid${i}_name`)
            localStorage.removeItem(`Gid${i}`)
            localStorage.removeItem(`good${i}`)
        }

    }


}
listening()

// 删除商品





let login_register = document.querySelector(".login_register").querySelector("li")
let userForm = document.querySelector(".banner").querySelector(".text").querySelector(".userForm").querySelectorAll("a")
if (localStorage.getItem("Token")) {


    console.log(userForm);
    // //访问他的孩子们
    (async function() {
        let user_form = await getToken("user/", localStorage.getItem("Token"))
            // console.log(user_form);
        userForm[0].innerHTML = ` 欢迎您， ${user_form.msg.Username}`

        userForm[1].remove()

        userForm[2].remove()

        userForm[0].nextSibling.nextSibling.textContent = "  ";
        userForm[0].style.top = "-35px"
        userForm[0].insertAdjacentHTML('afterend', '<input   type="button"  value="退出登录">')
        login_register.innerHTML = `${user_form.msg.Username}`
        login_register.style.color = "red";
        login_register.style.fontSize = "15px"


        let input = document.querySelector(".banner").querySelector(".text").querySelector(".userForm").querySelector("input") //获取推出登录节点
        input.addEventListener("click", () => {
            recover_no_login()
                // await out("logout")
        })
    })()
}



function recover_no_login() {
    let input = document.querySelector(".banner").querySelector(".text").querySelector(".userForm").querySelector("input")
    userForm[0].innerHTML = ` 欢迎来到 ，红岩网校`
    userForm[0].insertAdjacentHTML('afterend', '<a href="http://127.0.0.1:5500/login.html">登录</a>');
    userForm[0].insertAdjacentText("afterend", " | ")
    userForm[0].insertAdjacentHTML("afterend", '<a href="http: //127.0.0.1:5500/register-first.html">注册</a>')
    userForm[0].style.top = ""
    login_register.innerHTML = "免费注册"
    login_register.style.color = ""
    login_register.style.fontSize = ""
    console.log(userForm[2]);
    input.remove();
    (async() => {
        let res = await out("logout", localStorage.getItem("Token"))
        localStorage.removeItem("Token")
        localStorage.clear()
    })()

}
// 推出登录 函数 


//判断用户状态 0


//商品页面l
;
(async function() {
    let ul_list = document.querySelector(".kill_slide").querySelectorAll(".list")
    let a_list = document.querySelector(".good_list").querySelectorAll("a")
    let good_list = await getToken("show")
        // console.log(good_list);
    console.log(good_list.msg.All);


    function get_something(arr, get_something) {
        let arr0 = []
        let arr_introduce = []
        let arr_Gid = []
        let arr__price = []
        let arr_img = []
        let arr_name = []
        for (let i = 0; i < arr.length; i++) {
            let xxx = arr[i].Goods.filter((e, i, a) => {

                if (e[get_something]) arr0.push(e[get_something])
                if (e.Introduce) arr_introduce.push(e.Introduce)
                if (e.Gid) arr_Gid.push(e.Gid)
                if (e.Price) arr__price.push(e.Price)
                if (e.Url) arr_img.push(e.Url)
                if (e.Name) arr_name.push(e.Name)

                return arr0[i]
            })

        }

        arr_introduce.push(arr_introduce.splice(0, 1)[0])
        arr_introduce.push(arr_introduce.splice(0, 1)[0])
        let temp_introduce = arr_introduce.splice(3, 2)
        arr_introduce.splice(11, 0, temp_introduce[0], temp_introduce[1])

        arr_Gid.push(arr_Gid.splice(0, 1)[0])
        arr_Gid.push(arr_Gid.splice(0, 1)[0])
        let temp_Gid = arr_Gid.splice(3, 2)
        arr_Gid.splice(11, 0, temp_Gid[0], temp_Gid[1])

        arr__price.push(arr__price.splice(0, 1)[0])
        arr__price.push(arr__price.splice(0, 1)[0])
        let temp__price = arr__price.splice(3, 2)
        arr__price.splice(11, 0, temp__price[0], temp__price[1])


        arr_img.push(arr_img.splice(0, 1)[0])
        arr_img.push(arr_img.splice(0, 1)[0])
        let temp__src = arr_img.splice(3, 2)
        arr_img.splice(11, 0, temp__src[0], temp__src[1])

        arr_name.push(arr_name.splice(0, 1)[0])
        arr_name.push(arr_name.splice(0, 1)[0])
        let temp_name = arr_name.splice(3, 2)
        arr_name.splice(11, 0, temp_name[0], temp_name[1])
            //处理数据 按照正常顺序
        for (let i = 0; i < arr0.length; i++) {
            let good_list = {
                Gid: arr_Gid[i],
                introduce: arr_introduce[i],
                price: arr__price[i],
                src: arr_img[i],
                name: arr_name[i],
                number: 0

            }
            localStorage.setItem(`good_list${i}`, JSON.stringify(good_list))

        }

        arr0.push(arr0.splice(0, 1)[0])
        arr0.push(arr0.splice(0, 1)[0])
        return arr0
    }

    let good__img_list = get_something(good_list.msg.All, "Url"),
        good__Price_list = get_something(good_list.msg.All, "Price"),
        good__h4_list = get_something(good_list.msg.All, "Introduce"),
        good__h5_list = get_something(good_list.msg.All, "Name"),
        good__Gid_list = get_something(good_list.msg.All, "Gid")
    console.log(good__Gid_list);
    //console.log(good__img_list);
    Array.from(ul_list[0].children).forEach((e, index) => {
        e.children[0].src = good__img_list[0]
        e.children[1].innerHTML = good__h5_list[0]
        e.children[2].innerHTML = "￥" + good__Price_list[0]
        e.Gid = good__Gid_list[0]
        e.name = 0
        console.log(e.Gid);
    });
    Array.from(ul_list[1].children).forEach((e, index) => {
        e.children[0].src = good__img_list[1]
        e.children[1].innerHTML = good__h5_list[1]
        e.children[2].innerHTML = "￥" + good__Price_list[1]
        e.Gid = good__Gid_list[1]
        e.name = 1
        console.log(e.Gid);
    });
    Array.from(ul_list[2].children).forEach((e, index) => {
        e.children[0].src = good__img_list[2]
        e.children[1].innerHTML = good__h5_list[2]
        e.children[2].innerHTML = "￥" + good__Price_list[2]
        e.Gid = good__Gid_list[2]
        e.name = 2
        console.log(e.Gid);
    });

    // 给秒杀 赋值

    for (let i = 0; i < 6; i++) {
        a_list[i].children[0].src = good__img_list[i + 5]
        a_list[i].children[1].children[0].innerHTML = good__h5_list[i + 5]
        a_list[i].children[1].children[1].innerHTML = "￥" + good__Price_list[i + 5]
        a_list[i].Gid = good__Gid_list[i + 5]
        a_list[i].name = i + 3
        console.log(a_list[i].Gid);
    }
    // for (let i = 7; i < 8; i++) {
    //     a_list[i].children[0].src = good__img_list[i + 5]
    //     a_list[i].children[1].children[0].innerHTML = good__h5_list[i + 5]
    //     a_list[i].children[1].children[1].innerHTML = "￥" + good__Price_list[i + 5]

    // }

    a_list[6].children[0].src = good__img_list[11]
    a_list[6].children[1].children[0].innerHTML = good__h5_list[11]
    a_list[6].children[1].children[1].innerHTML = "￥" + good__Price_list[11]
    a_list[6].Gid = good__Gid_list[11]
    a_list[6].name = 9
    console.log(a_list[6].Gid);
    for (let i = 7; i < 15; i++) {
        a_list[i].children[0].src = good__img_list[12]
        a_list[i].children[1].children[0].innerHTML = good__h5_list[12]
        a_list[i].children[1].children[1].innerHTML = "￥" + good__Price_list[12]
        a_list[i].Gid = good__Gid_list[12]
        a_list[i].name = 10
        console.log(a_list[i].Gid);
    }
    a_list[15].children[0].src = good__img_list[4]
    a_list[15].children[1].children[0].innerHTML = good__h5_list[4]
    a_list[15].children[1].children[1].innerHTML = "￥" + good__Price_list[4]
    a_list[15].Gid = good__Gid_list[4]
    a_list[15].name = 11
    console.log(a_list[15].Gid);
    for (let i = 16; i < 24; i++) {
        a_list[i].children[0].src = good__img_list[3]
        a_list[i].children[1].children[0].innerHTML = good__h5_list[3]
        a_list[i].children[1].children[1].innerHTML = "￥" + good__Price_list[3]
        a_list[i].Gid = good__Gid_list[3]
        a_list[i].name = 12
        console.log(a_list[i].Gid);

    }
    a_list[24].children[0].src = good__img_list[13]
    a_list[24].children[1].children[0].innerHTML = good__h5_list[13]
    a_list[24].children[1].children[1].innerHTML = "￥" + good__Price_list[13]
    a_list[24].Gid = good__Gid_list[13]
    a_list[24].name = 13
    console.log(a_list[24].Gid);
    for (let i = 25; i < 33; i++) {
        a_list[i].children[0].src = good__img_list[14]
        a_list[i].children[1].children[0].innerHTML = good__h5_list[14]
        a_list[i].children[1].children[1].innerHTML = "￥" + good__Price_list[14]
        a_list[i].Gid = good__Gid_list[14]
        a_list[i].name = 14
        console.log(a_list[i].Gid);
    }


    // Array.from(a_list[0]).forEach(e => {
    //         console.log(good__img_list[1]);
    //         e.children[0].src = good__img_list[1]
    //         console.log(e.children[0].src);
    //     })
    // console.log(a_list[0].children[0]);
    // console.log(a_list[0].children[0].src);

})()




// 商品绑定 跳转 存入  本地数据

let everyDay_sell = document.querySelector(".good_list")

console.log(everyDay_sell.children);
everyDay_sell.addEventListener("click", (e) => {

        if (e.target.tagName == "H4" || e.target.tagName == "IMG" || e.target.tagName == "H5") {

            if (e.target.tagName == "H4" || e.target.tagName == "H5") {
                console.log(e.target.parentNode.parentNode);
                console.log(e.target.parentNode.parentNode.children[0].src);
                console.log(e.target.parentNode.parentNode.children[1].children[0].innerHTML);
                console.log(e.target.parentNode.parentNode.children[1].children[1].innerHTML);

                localStorage.setItem(`Gid${e.target.parentNode.parentNode.name}_src`, JSON.stringify(e.target.parentNode.parentNode.children[0].src))
                localStorage.setItem(`Gid${e.target.parentNode.parentNode.name}_name`, JSON.stringify(e.target.parentNode.parentNode.children[1].children[0].innerHTML))
                localStorage.setItem(`Gid${e.target.parentNode.parentNode.name}_price`, JSON.stringify(e.target.parentNode.parentNode.children[1].children[1].innerHTML))
                localStorage.setItem(`Gid${e.target.parentNode.parentNode.name}_id`, JSON.stringify(e.target.parentNode.parentNode.name))
                localStorage.setItem(`Gid${e.target.parentNode.parentNode.name}`, JSON.stringify(e.target.parentNode.parentNode.Gid))
                    // localStorage.setItem("", )
            }

            if (e.target.tagName == "IMG") {
                console.log(e.target.parentNode.name);
                console.log(e.target.parentNode.children[0].src);
                console.log(e.target.parentNode.children[1].children[0].innerHTML);
                console.log(e.target.parentNode.children[1].children[1].innerHTML);
                localStorage.setItem(`Gid${e.target.parentNode.name}_src`, JSON.stringify(e.target.parentNode.children[0].src))
                localStorage.setItem(`Gid${e.target.parentNode.name}_name`, JSON.stringify(e.target.parentNode.children[1].children[0].innerHTML))
                localStorage.setItem(`Gid${e.target.parentNode.name}_price`, JSON.stringify(e.target.parentNode.children[1].children[1].innerHTML))
                localStorage.setItem(`Gid${e.target.parentNode.name}_id`, JSON.stringify(e.target.parentNode.name))
                localStorage.setItem(`Gid${e.target.parentNode.name}`, JSON.stringify(e.target.parentNode.Gid))
            }



            // console.log(e.target.parentNode.children[0].src);
            // console.log(e.target.parentNode.children[1].children);
            //, e.target.parentNodee.children[2].innerHTML
            // e.children[0].src 
            // e.children[1].innerHTML 
            // e.children[2].innerHTML 

        };
    }, false)
    //商品页面











// second kill  部分  倒计时部分 
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





(async() => {
    let res = await get_news("news")
    console.log(res.data[0].title);
    let news = document.querySelector(".news").querySelectorAll("i");
    [...news].forEach((e, i) => {
        e.innerText = res.data[i].title
    })
})()



function shop() {
    let myshop = document.querySelector(".myshop")
    console.log(myshop);
    myshop.innerText = localStorage.getItem("shopping_number")

}
shop()