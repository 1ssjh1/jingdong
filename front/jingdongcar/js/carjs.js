import { postFormDate, postForm, getToken, car_put, postddd } from "../../module/module.js"


let listening_code = true

function listening() {
    for (let i = 0; i < 15; i++) {
        if (listening_code) {

            localStorage.removeItem(`good${i}`)
        }
    }
}
listening()


// 右上角  的购物车 小按钮

function cart_nav() {
    let car_number = document.querySelector(".car_number")
    let temp_number = localStorage.getItem("shopping_number") || "0"
        // console.log(car_number.innerText);


    car_number.innerText = temp_number
}
cart_nav()

// 右上角  的购物车 小按钮


function sum_fn() {
    let total = document.querySelector(".total")

    let sum = document.querySelectorAll(".sum")
    let total_price = 0
    Array.from(sum).forEach(e => {
        if (e.checked == true) {


            console.log(e.parentElement.parentElement.lastElementChild.previousElementSibling.children[0].innerText);
            total_price += parseFloat(e.parentElement.parentElement.lastElementChild.previousElementSibling.children[0].innerText, 2)
            console.log(total_price);
        }

    })
    total.innerText = total_price

}


// buy——buy
console.log(localStorage.getItem("shopping_number"));
let btn_buy = document.querySelector(".buy")
btn_buy.addEventListener("click", (e) => {
    e.preventDefault()

    let list_goods = document.querySelectorAll(".list_goods")
    console.log(list_goods);
    let form = new FormData()
    let length = localStorage.getItem("shopping_number")
    console.log(form.get("chart_id"));
    [...list_goods].forEach((e) => {
        if (e.children[0].children[0].checked == true) {
            form.append("chart_id", e.lastElementChild.previousElementSibling.previousElementSibling.children[0].children[0].getAttribute("chart_id") / 1)
            e.remove()
            length--
        }
        console.log(form.getAll("chart_id"));

    })

    buy(form)


    console.log(localStorage.getItem("shopping_number"));
    localStorage.setItem("shopping_number", length)
    cart_nav()
})

function buy(form) {


    (async() => {
        let res = await postddd("shop/order", localStorage.getItem("Token"), form)
        console.log(res);

        if (res.state == true) {
            alert("正在跳转中...")
            setTimeout(() => {
                window.location.href = "http://127.0.0.1:5500/user.html"

            }, 2000);

        } else {
            let false_msg = document.querySelector(".false_msg")
            false_msg.innerText = res.msg
        }
    })()




}


// buy——buy









function Car() {


    this.object = {}
}
let car = new Car()




Car.prototype.creat_begin = function() {

        (async() => {
            let res = await getToken("shop/chart", localStorage.getItem("Token"))
            console.log(res.msg.ChartList);


            let temp_goods_number = 0
            let car_top = document.querySelector('.car_top').children[0].children[0].children[0]
            let car_center = document.querySelector(".car_center")
            for (let x = 0; x < 14; x++) {
                for (let y = 0; y < res.msg.ChartList.length; y++) {

                    if (res.msg.ChartList[y].Gid == JSON.parse(localStorage.getItem(`good_list${x}`)).Gid) {
                        localStorage.setItem(`chart_id${y}`, res.msg.ChartList[y].ChartId)
                        let Chart_id = localStorage.getItem(`chart_id${y}`)
                        console.log(Chart_id);
                        let temp = JSON.parse(localStorage.getItem(`good_list${x}`))
                        temp.number = res.msg.ChartList[y].Count


                        console.log(temp.price);
                        let temp_introduce = temp.introduce

                        let small_price = parseFloat(temp.price / 1) * (temp.number / 1)
                            // let this_price = parseFloat((temp.price.toString().split("￥")[1] / 1))
                            // let small_price = parseFloat((temp.price.toString().split("￥")[1] / 1)) * (temp.number / 1)
                        car_center.insertAdjacentHTML("afterbegin", `  <ul class="list_goods">
                            <li><input type="checkbox" class="button sum"> </li>
                            <li><img src="${temp.src}" alt="" class="goods_img" ></li>
                            <li><i>${temp.name}</i></li>
                            <li><i>${temp_introduce} </i></li>
        
                            <li> ￥<i>${temp.price}</i></li>
                            <li>
                             <div class="choose_amount">
                    <input type="text" value="${temp.number}" class="text" gid="${temp.Gid}" Chart_id="${Chart_id}">
                    <a href="javascript:;" class="add">+</a>
                    <a href="javascript:;" class="reduce">-</a>
                        </div>
                    </li>
                    <li>￥<i class="small_total">${small_price}</i></li>
                    <li><a href="javascript:;" class="delete">删除</a></li>
                    </ul>`)

                        temp_goods_number++

                        car_top.innerHTML = temp_goods_number
                    }

                }



            }

            localStorage.setItem("shopping_number", temp_goods_number)
            cart_nav()
            car.select()
            car.price_number()
            dele()









        })()










    }
    // class="all_select"


Car.prototype.select = function() {
    let all_select = document.querySelectorAll(".all_select")
    let button = document.querySelectorAll(".button")
    console.log(button);
    let car_center = document.querySelector(".car_center")
    let input_list = document.querySelector(".car_center").querySelectorAll("input")
    let select_number = 0
    let goods_number = document.querySelector(".goods_number")
    let key = true //全选🔒
    console.log(all_select[0]);
    all_select[0].addEventListener("click", () => {
        Array.from(button).forEach((e) => {
            e.checked = true
            console.log(button);
            all_select[1].checked = false
        })

        goods_number.innerText = input_list.length
        sum_fn()

    })

    all_select[1].addEventListener('click', () => {
        Array.from(button).forEach((e) => {
            e.checked = false
            all_select[0].checked = false
        })
        goods_number.innerText = 0
        sum_fn()

    })


    car_center.addEventListener('click', (e) => {

        if (e.target.tagName.toLowerCase() == "input") {
            all_select[0].checked = false;
            all_select[1].checked = false;
            [...input_list].forEach((e) => {
                if (e.checked == true)
                    select_number++;
                if (e.checked == false)
                    key = false
                else
                    key = true
                all_select[0].checked = key
            })

            goods_number.innerText = select_number
            select_number = 0
            sum_fn()

        }

    }, true)

    // Array.prototype.slice.call(input_list, 0).forEach(element => {

    // });


}


Car.prototype.price_number = function() {

    let small_total = document.querySelectorAll(".small_total")



    let price_change = document.querySelectorAll(".choose_amount");

    [...price_change].forEach(e => {
        e.addEventListener("click", (e) => {

            if (e.target.parentNode.children[0].value >= 20) {
                alert("太多了太多了！！")
                e.target.parentNode.children[0].value--
                    return
            }
            if (1 >= e.target.parentNode.children[0].value && e.target.className != "add") {
                alert("不能再少了")
                e.target.parentNode.children[0].value++

            }


            if (e.target.className == "add") {
                e.target.parentNode.children[0].value++
                    let chart_id = e.target.parentNode.children[0].getAttribute("Chart_id") / 1,
                        Count = e.target.parentNode.children[0].value / 1

                ;
                (async() => {

                    let add_put = {
                        chart_id: chart_id,
                        Count: Count
                    }
                    console.log(add_put);
                    let res = await car_put("shop/chart", localStorage.getItem("Token"), add_put)
                    console.log(res);
                })()


            }
            if (e.target.className == "reduce") {
                e.target.parentNode.children[0].value--
                    let chart_id = e.target.parentNode.children[0].getAttribute("Chart_id") / 1,
                        Count = e.target.parentNode.children[0].value / 1

                ;
                (async() => {
                    let add_put = {
                        chart_id: chart_id,
                        Count: Count
                    }
                    console.log(add_put);
                    let res = await car_put("shop/chart", localStorage.getItem("Token"), add_put)
                    console.log(res);
                })()


            }
            console.log(e.target.parentNode.parentNode.nextElementSibling.children[0].innerText)

            //修改前面的并列处理 
            e.target.parentNode.parentNode.nextElementSibling.children[0].innerText = (e.target.parentNode.children[0].value) / 1 * e.target.parentNode.parentNode.previousElementSibling.children[0].innerText


            sum_fn()

        }, false)

        // function sum_fn() {
        //     let total = document.querySelector(".total")

        //     let sum = document.querySelectorAll(".sum")
        //     let total_price = 0
        //     Array.from(sum).forEach(e => {
        //         if (e.checked == true) {


        //             console.log(e.parentElement.parentElement.lastElementChild.previousElementSibling.children[0].innerText);
        //             total_price += parseFloat(e.parentElement.parentElement.lastElementChild.previousElementSibling.children[0].innerText, 2)
        //             console.log(total_price);
        //         }

        //     })
        //     total.innerText = total_price

        // }
        sum_fn()
    });






}





function dele() {
    let delete_btn = document.querySelectorAll(".delete")


    Array.from(delete_btn).forEach((e) => {
        e.addEventListener("click", function(e) {



            let temp_shopping_number = localStorage.getItem("shopping_number")
            console.log(temp_shopping_number);
            localStorage.setItem("shopping_number", (temp_shopping_number / 1) - 1)
            cart_nav()
            this.parentNode.parentNode.remove()
            let chart_id = e.target.parentNode.previousElementSibling.previousElementSibling.children[0].children[0].getAttribute("Chart_id") / 1;
            console.log(chart_id);;
            (async() => {
                let add_put = {
                    chart_id: chart_id,
                    Count: 0
                }
                console.log(add_put);
                let res = await car_put("shop/chart", localStorage.getItem("Token"), add_put)
                console.log(res);
            })()




        })
    })

}
dele()





car.creat_begin()
car.select()
car.price_number()