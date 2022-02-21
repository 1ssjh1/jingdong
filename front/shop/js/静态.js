import { postForm, getToken, getToken_msg } from "../../module/module.js"

let good_img = document.querySelector(".big_img") //获取图片

let name = document.querySelector(".name") //获取相关信息

let listName = document.querySelector(".listName")

let Price = document.querySelector(".price")
let listening_code = false






function listening() {
    for (let i = 0; i < 15; i++) {
        if (listening_code) {
            localStorage.removeItem(`Gid${i}_src`)
            localStorage.removeItem(`Gid${i}_price`)
            localStorage.removeItem(`Gid${i}_name`)
            localStorage.removeItem(`Gid${i}_id`)
            localStorage.removeItem(`Gid${i}`)

        }

    }


}
listening()



for (let i = 0; i < 15; i++) {

    if (localStorage.getItem(`Gid${i}_src`)) {

        good_img.children[0].src = JSON.parse(localStorage.getItem(`Gid${i}_src`));
        Price.innerHTML = JSON.parse(localStorage.getItem(`Gid${i}_price`));
        name.children[1].innerText = ">" + JSON.parse(localStorage.getItem(`Gid${i}_name`));
        listName.innerHTML = "[爆款！ 爆款！]" + JSON.parse(localStorage.getItem(`good_list${i}`)).introduce;

        console.log(JSON.parse(localStorage.getItem(`Gid${i}_price`)));


        (async() => {
            let gid = JSON.parse(localStorage.getItem(`Gid${i}`))
            let res = await getToken_msg("shop/commit?gid=", localStorage.getItem("Token"), gid)
            let user_assess = document.querySelector(".user_assess")
            let all_commit = document.querySelector(".all_commit")
            let good_commit = document.querySelector(".good_commit")
            console.log(res);
            console.log(res.commit.onecomit[0].commit);

            for (let i = 0; i < res.commit.onecomit.length; i++) {
                all_commit.innerText++
                    good_commit.innerText++
                    user_assess.insertAdjacentHTML("afterbegin", ` 
             <li class="user">
             <div class="pic">
             <a href=""><img src="${res.commit.onecomit[i].url}" alt=""></a>
             </div>
             <div class="user_review"> <span>${res.commit.onecomit[i].commit}</span></div>
             </li>`)
                console.log(res);
            }
        })()



        // localStorage.removeItem(`Gid${i}_src`)
        // localStorage.removeItem(`Gid${i}_price`)
        // localStorage.removeItem(`Gid${i}_name`)

        // window.pushState = function() {
        //     alert(111)
        //     localStorage.removeItem(`Gid${i}_src`)
        //     localStorage.removeItem(`Gid${i}_price`)
        //     localStorage.removeItem(`Gid${i}_name`)

        // }

        // if (localStorage.getItem(`good${i}`)) {
        //     let temporary_good = JSON.parse(localStorage.getItem(`good${i}`))
        //     if (temporary_good.number == 0) {

        //     }



        // }
        if (!localStorage.getItem(`good${i}`)) {
            let to_good = {
                    name: JSON.parse(localStorage.getItem(`Gid${i}_name`)),
                    price: JSON.parse(localStorage.getItem(`Gid${i}_price`)),
                    number: 0,
                    src: JSON.parse(localStorage.getItem(`Gid${i}_src`)),
                    introduce: "[爆款！ 爆款！]" + JSON.parse(localStorage.getItem(`good_list${i}`)).introduce

                }
                //＋id
            localStorage.setItem(`good${i}`, JSON.stringify(to_good))
        }
    }
    console.log(localStorage.getItem(`Gid${i}_src`));






}




//点击加量购买 
let choose_amount = document.querySelector(".choose_amount")
choose_amount.addEventListener("click", (e) => {
    if (e.target.parentNode.children[0].value >= 20)

    {
        alert("太多了太多了！！")
        e.target.parentNode.children[0].value--
            return
    }

    if (0 >= e.target.parentNode.children[0].value) {
        alert("不能再少了")
        e.target.parentNode.children[0].value++
            return
    }

    if (e.target.className == "add") {
        e.target.parentNode.children[0].value++

    }
    if (e.target.className == "reduce") {
        e.target.parentNode.children[0].value--

    }


})




//加入购物车  
let add_car = document.querySelector(".addCar")
add_car.addEventListener("click", (e) => {
    e.preventDefault()
    console.log(choose_amount.children[0].value);


    for (let i = 0; i < 15; i++) {

        if (localStorage.getItem(`Gid${i}_src`)) {

            let new_good = JSON.parse(localStorage.getItem(`good${i}`))
                // console.log(new_good);


            // console.log(localStorage.getItem(`Gid${i}`),
            localStorage.setItem(`good${i}`, JSON.stringify(new_good))
                // choose_amount.children[0].value * 1

            ;
            (async() => {
                let gid = localStorage.getItem(`Gid${i}`) * 1,
                    count = choose_amount.children[0].value * 1
                console.log(gid, count);
                let data = {
                    gid: gid,

                    count: count
                }


                let res = await postForm("shop/chart", data, localStorage.getItem("Token"))
                console.log(res.msg);
                alert(res.msg)
            })()
        }

    }
    cart_nav()




    // window.location.href = "http://127.0.0.1:5500/car.html"
})









// 右上角  的购物车 小按钮

function cart_nav() {
    let car_number = document.querySelector(".car_number")
    let temp_number = localStorage.getItem("shopping_number") || "0"
        // console.log(car_number.innerText);

    temp_number++
    car_number.innerText = temp_number
    temp_number--
}
cart_nav()

// 右上角  的购物车 小按钮




function good_select(choose) {
    let thing = document.querySelector(`.${choose}`).children[1]
    thing.addEventListener("click", function(e) {
        if (e.target.tagName.toLowerCase() == "a") {
            for (let i = 0; i < thing.children.length; i++) {
                thing.children[i].className = ""
                if (thing.children[i] == e.target) {

                    e.target.className = "current"
                }
            }
        }
    })
}
good_select("choose_color")
good_select("choose_version")
good_select("choose_type")
good_select("product_msg .choose_version_b")