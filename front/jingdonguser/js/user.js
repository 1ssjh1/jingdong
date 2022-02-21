import { postFormDate, postForm, car_put, getToken, postToken, post_comment, delete_comment, post_recharge } from "../../module/module.js"

function shift() {

    let shift = document.querySelector(".shift"),
        shift_list = document.querySelector(".shift").querySelectorAll("li"),
        div_list = document.querySelector(".orders_nav").querySelectorAll(".orders_nav>div")
    console.log(div_list);
    shift.addEventListener("click", (e) => {
        let index = e.target.parentElement.id;

        if (index == 3) {
            let finish_ul = e.target.parentElement.parentElement.parentElement.lastElementChild
            Array.from(finish_ul.children).map((e) => {
                e.remove()
            });
            (async() => {
                let new_user = await getToken("user/ ", localStorage.getItem("Token"))

                finish(new_user.msg.Category[2].Order)
                comments()
                delete_select()
            })()


        }




        [...shift_list].forEach((e, i) => {

            e.className = ""
            if (i == index) {
                e.className = "focus_color"
            }

        });


        if (e.target.tagName.toLowerCase() == "h4") {

            Array.from(div_list).forEach((e, i) => {
                e.style.display = "none"
                e.className
                if (i == index) {
                    e.style.display = "block"
                }

            })

        }



    }, false)


}
shift()




function creat_userForm() {
    let img_name = document.querySelector(".avatar"),
        balance = document.querySelector(".balance"),
        preferential = document.querySelector(".preferential");
    (async() => {
        let user_form = await getToken("user/ ", localStorage.getItem("Token"))
        console.log(user_form.msg);
        try {
            img_name.children[0].src = user_form.msg.ImageUrl
        } catch (err) {
            console.log("img not find")
        }

        img_name.children[1].innerHTML = user_form.msg.Username,
            balance.innerHTML = "￥" + user_form.msg.Balance
        preferential.innerHTML = user_form.msg.Uid

        let key = true

        function hid() {

            let hid = document.querySelector(".hid"),
                money = document.querySelector("#money"),
                temp_money = money.innerText
            hid.addEventListener("click", () => {
                if (key) {
                    money.innerText = "* * * "
                    key = false
                } else {
                    key = true;
                    money.innerText = temp_money
                }


            })

        }
        hid()







        receipt(user_form.msg.Category[0].Order)


        appraise(user_form.msg.Category[1].Order)


        finish(user_form.msg.Category[2].Order)

        goods()
        touxiang()





        shop()
        make_sure()




    })()
}
creat_userForm()

function goods() {
    (async() => {
        let payment = document.querySelector(".payment")
        let res = await getToken("shop/chart", localStorage.getItem("Token"))
        console.log(res.msg.ChartList);
        for (let x = 0; x < 15; x++) {

            for (let y = 0; y < res.msg.ChartList.length; y++) {

                if (res.msg.ChartList[y].Gid == JSON.parse(localStorage.getItem(`good_list${x}`)).Gid) {
                    let temp = JSON.parse(localStorage.getItem(`good_list${x}`))
                    let temp_count = res.msg.ChartList[y]

                    payment.insertAdjacentHTML("afterbegin",
                        ` <ul class="list_goods">
                    <li><img src="${temp.src}" alt=""></li>
                    <li> <i>${temp.introduce}</i> </li>
                    <li >
                     
                    </li>
                    <li> <i>${temp_count.Count}</i></li>
                    <li>
                        <div class="choose_amount">

                            <a href="javascript:;" class="reduce">$${temp.price}</a>
                        </div>
                    </li>
                    <li><i>等待付款</i></li>
                    <li class="payment_button"><input type="button" value="去付款" id="1"></li>
                </ul>`



                    )
                }




            }
        }
    })()
}





function receipt(receipt) {

    let receipted = document.querySelector(".receipt")

    for (let x = 0; x < 15; x++) {

        for (let y = 0; y < receipt.length; y++) {


            if (receipt[y].Gid == JSON.parse(localStorage.getItem(`good_list${x}`)).Gid)

            {
                let temp = JSON.parse(localStorage.getItem(`good_list${x}`))

                receipted.insertAdjacentHTML("afterbegin",
                    ` <ul class="list_goods">
                    <li><img src="${temp.src}" alt=""></li>
                    <li> <i>${temp.introduce}</i> </li>
                    <li >
                     
                    </li>
                    <li> <i>${receipt[y].Count}</i></li>
                    <li>
                        <div class="choose_amount">

                            <a href="javascript:;" class="reduce">$${temp.price}</a>
                        </div>
                    </li>
                    <li><i>已支付</i></li>
                    <li class="payment_button"><input type="button" value="等待发货" id="${receipt[y].Oid}"></li>
                </ul>`



                )




            }

        }
    }

}


function appraise(appraise) {

    let appraised = document.querySelector(".appraise")

    for (let x = 0; x < 15; x++) {

        for (let y = 0; y < appraise.length; y++) {


            if (appraise[y].Gid == JSON.parse(localStorage.getItem(`good_list${x}`)).Gid)

            {
                let temp = JSON.parse(localStorage.getItem(`good_list${x}`))

                appraised.insertAdjacentHTML("afterbegin",
                    ` <ul class="list_goods">
                    <li><img src="${temp.src}" alt=""></li>
                    <li> <i>${temp.introduce}</i> </li>
                    <li >
                     
                    </li>
                    <li> <i>${appraise[y].Count}</i></li>
                    <li>
                        <div class="choose_amount">

                            <a href="javascript:;" class="reduce">$${temp.price}</a>
                        </div>
                    </li>
                    <li><i>已支付</i></li>
                    <li class="payment_button"><input type="button" value="确认收货" id="${appraise[y].Oid}" class="make_sure"></li>
                </ul>`



                )




            }

        }
    }

}

function finish(finish) {

    let finished = document.querySelector(".finish")

    for (let x = 0; x < 15; x++) {

        for (let y = 0; y < finish.length; y++) {


            if (finish[y].Gid == JSON.parse(localStorage.getItem(`good_list${x}`)).Gid)

            {
                let temp = JSON.parse(localStorage.getItem(`good_list${x}`))

                finished.insertAdjacentHTML("afterbegin",
                    ` <ul class="list_goods">
                    <li><img src="${temp.src}" alt=""></li>
                    <li> <i>${temp.introduce}</i> </li>
                    <li class="user_appraise">
                    <textarea cols="15" rows="3" style="resize:none;" placeholder="写出你的评价"></textarea>
                    <input type="button" value="提交" class="submit" id="${finish[y].Oid}" ><i></i>
                    </li>
                    <li> <i>${finish[y].Count}</i></li>
                    <li>
                        <div class="choose_amount">

                            <a href="javascript:;" class="reduce">$${temp.price}</a>
                        </div>
                    </li>
                    <li><i>已支付</i></li>
                    <li class="finish_button payment_button"><input type="button" value="删除订单" id="${finish[y].Oid}" ></li>
                </ul>`



                )




            }

        }
    }

}



function make_sure() {
    let make_sure = document.querySelector(".appraise").querySelectorAll(".make_sure")
    Array.from(make_sure).forEach((e) => {
        e.addEventListener("click", function(e) {

            if (confirm("是否确认收货")) {
                e.target.parentElement.parentElement.remove();
                (async() => {
                    let oid = this.id / 1
                    let make_sure_btn = {
                        oid: oid,
                    }
                    let res = await car_put("user/order", localStorage.getItem("Token"), make_sure_btn)
                    console.log(res);
                })()
            }


        })

    })
}



function comments() {
    let btn_list = document.querySelectorAll(".submit")
        // btn_list = document.querySelectorAll
    Array.from(btn_list).forEach((e) => {

        e.addEventListener("click", (element) => {

            // console.log(element.target.previousElementSibling.value);
            if (!(/^[\u4e00-\u9fa5]{5,15}$|^[\dA-Za-z_]{6,30}$/).test(element.target.previousElementSibling.value)) {

                element.target.nextElementSibling.innerText = "过长过短都不行！"
                element.target.nextElementSibling.className = "color_red"
            } else {
                let commit = element.target.previousElementSibling.value;
                let oid = element.target.id / 1
                element.target.previousElementSibling.readonly = "true"
                console.log(element.target.previousElementSibling);
                element.target.nextElementSibling.innerText = "OK"
                element.target.nextElementSibling.className = "color_green"
                console.log(element.target);
                element.target.remove();

                (async() => {
                    let formData = new FormData()
                    formData.append("oid", oid)
                    formData.append("commit", commit)
                    let res = await post_comment("user/commit", localStorage.getItem("Token"), formData)
                    if (res.state == false) {
                        alert(res.msg)
                    }

                })()
            }
        })
    })



}



function delete_select() {
    let finish_button = document.querySelectorAll(".finish_button")
    console.log(finish_button);
    let arr = Array.apply(null, finish_button)
    arr.forEach(e => {

        e.addEventListener("click", (element) => {

            if (confirm('确定删除？订单将不再恢复')) {
                element.target.parentElement.parentElement.remove()
                let oid = element.target.id / 1;


                (async() => {
                    let formData = new FormData()
                    formData.append("oid", oid)
                    let res = await delete_comment("user/order", localStorage.getItem("Token"), formData)
                    console.log(res);
                })()
            }




        })
    })
}




function shop() {
    let myshop = document.querySelector(".myshop")
    console.log(myshop);
    myshop.innerText = localStorage.getItem("shopping_number")

}

function touxiang() {

    let formElem = document.querySelector("#formElem")

    formElem.onsubmit = async(e) => {
        e.preventDefault();

        let response = await fetch('https://sanser.ltd/user/image', {
            method: 'PUT',
            headers: {
                'Authorization': `${localStorage.getItem("Token")}`,
                'Content-Type': 'multipart/form-data'
                    // 'Content-Type': `application/json`
                    // 'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: new FormData(formElem)
        });

        let result = await response.json();
        console.log(result);

    }
}



function recharge() {
    let recharge = document.querySelector(".recharge")



    recharge.addEventListener("click", (e) => {
        let money = prompt("请问你充值多少人名币").replace(/[^0-9]/ig, "") / 1

        ;
        (async() => {
            let form = new FormData()
            form.append("balance", money)
            let res = await post_recharge("user/balance", localStorage.getItem("Token"), form)

            if (res.state == true) {
                location.reload();
            }
        })()


    })
}
recharge()