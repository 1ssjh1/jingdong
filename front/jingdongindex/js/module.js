export { lunbo0, lunbo1, get, get_news, postForm, getMessage, postFormDate, getToken, out, seconds_kill }




function seconds_kill() {
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
}

function lunbo0() {
    let key = true //节流
    let index = 0
    let arr = []
    let cur_ul = document.querySelector(".banner_slide ")
    let lilist = document.querySelector(".banner_slide ").querySelectorAll('li')
    let timer = setInterval(go_next, 2000)


    for (let i = 0; i < 5; i++) {
        cur_ul.children[i].children[0].src = "./jingdongindex//imgs/" + (i + 1) + ".jpg";
        cur_ul.children[i].children[0].style.width = "590px";
        cur_ul.children[i].children[0].style.height = "470px";
        arr.push(cur_ul.children[i])
        arr[i].style.left = "0px"
        cur_ul.children[i].children[0].addEventListener("mouseenter", () => {
            clearTimeout(timer)
        })
        cur_ul.children[i].children[0].addEventListener("mouseleave", () => {
            timer = setInterval(go_next, 2000)
        })
    }
    cur_ul.insertAdjacentHTML("afterbegin", '<span class="iconfont pre" style="z-index:50">&#xe660;</span>')
    cur_ul.insertAdjacentHTML("beforeend", '<span class="iconfont next" style="z-index:50">&#xe65f;</span>')
    let pre = document.querySelector(".banner_slide ").querySelectorAll("span")[0]
    let next = document.querySelector(".banner_slide ").querySelectorAll("span")[1]
    for (let i = 0; i < arr.length; i++) {
        cur_ul.insertAdjacentHTML("beforeend", `<div style="left:${(i+1)*25}px;" id="${i}"></div>`);
    }
    let divlist = document.querySelector(".banner_slide ").querySelectorAll('div')


    go_index(index)
    for (let i = 0; i < divlist.length; i++) {
        divlist[i].addEventListener("click", function() {
            index = this.id
            go_index(index)
        })

    }

    function clear() {
        for (let i = 0; i < lilist.length; i++) {
            lilist[i].className = ""
            divlist[i].className = ""
        }
    }

    function go_index(index) {

        clear();
        lilist[index].className = "active";
        divlist[index].className = "point";
    }

    function go_next() {
        if (!key) return
        if (index < 4) index++
            else index = 0
        clear()
        go_index(index)
        key = false
        setTimeout(() => {
            key = true
        }, 1000);
    }

    function go_pre() {
        if (!key) return
        if (index > 0) index--
            else index = 4
        clear()
        go_index(index)
        key = false
        setTimeout(() => {
            key = true
        }, 1000);
    }

    next.addEventListener("click", () => {
        go_next()
    })
    pre.addEventListener("click", () => {
        go_pre()
    })
}

function lunbo1() {
    let index = 0
    let ul_list = document.querySelector(".kill_slide").querySelector(".warp").querySelectorAll("ul")
        // ul_list[0].style.left = "-1900px"
    let warp = document.querySelector(".kill_slide").querySelector(".warp")
    let pointList = document.querySelector(".kill_slide").querySelector(".point").querySelectorAll("li")
    let pre = document.querySelector(".kill_slide").querySelectorAll("span")[0]
    let next = document.querySelector(".kill_slide").querySelectorAll("span")[1]
    let clone = ul_list[0].cloneNode(true)
    ul_list[ul_list.length - 1].after(clone)
        //想前滚动函数 
    pre.addEventListener("click", () => {
            get_pre()
        })
        //向后滚动函数
    next.addEventListener("click", () => {
        get_next()
    })


    // 节流 
    let lock = true;
    //前进函数
    function get_next() {
        if (!lock) return
        if (index < 2) {
            warp.style.transition = "1s ease";
            index++;
            warp.style.left = -index * 800 + "px"
        } else {
            index++
            warp.style.left = -index * 800 + "px";
            index = 0
            setTimeout(() => {
                warp.style.left = -index * 800 + "px";
                warp.style.transition = "none";
            }, 800)
        }
        get_active()
        lock = false
        setTimeout(() => {
            lock = true
        }, 1000);

    }
    // 退后函数 



    function get_pre() {
        if (!lock) return;
        if (index > 0) {
            index--;
            warp.style.left = -index * 800 + "px"

        } else {
            warp.style.transition = "none";
            warp.style.left = -ul_list.length * 800 + "px"
            index = 2

            setTimeout(() => {
                warp.style.left = -index * 800 + "px"
                warp.style.transition = "1s ease";
            }, 0);
        }
        get_active()


        lock = false
        setTimeout(() => {
            lock = true
        }, 1000);
    }
    //小圆点部分高亮
    function get_active() {
        pointList.forEach((item, i, arr) => {
            if (arr[i].getAttribute("name") == index) pointList[i].className = "active"
            else pointList[i].className = ""
        });
    }
    //初始化
    get_active()
        // pointList.__proto__ = Array.prototype
        //绑定name属性
    pointList.forEach((item, index, arr) => {
        item.abc = index;


        item.style.left = index * 50 + 200 + "px"

    });
    //小圆点高亮
    pointList[0].parentElement.addEventListener("click", (e) => {
            if (e.target.nodeName.toLowerCase() == "li") {
                index = Number(e.target.getAttribute("name"));;
                get_active()
                warp.style.transition = "1s ease";
                warp.style.left = index * -800 + "px";
            }
        })
        //自动轮播
    let timer = setInterval(get_next, 3000);
    // 鼠标悬停 保持 不动
    warp.addEventListener('mouseenter', () => {
        clearTimeout(timer)
    })
    warp.addEventListener('mouseleave', () => {
        timer = setInterval(get_next, 3000);
    })
}




const Base_Url = "https://sanser.ltd/"


// -------------------------------------
// -------------------------------------
async function get_news(url) {
    const require = await fetch(`${Base_Url}${url}`, {

        headers: {
            'Content-Type': `application/json`,
        }

    })
    return await require.json()
}

async function postForm(url, data) {
    const require = await fetch(`${Base_Url}${url}`, {
        method: "POST",
        headers: {
            'Content-Type': `application/json`,
            //  'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: JSON.stringify(data)
    })

    return await require.json()
}


// -------------------------------------
async function postFormDate(url, Form) {
    const require = await fetch(`${Base_Url}${url}`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: Form
    })

    return await require.json()
}

// -------------------------------------

async function get(url) {
    const require = await fetch(`${Base_Url}${url}`, {})

    return await require.json()
}
// -------------------------------------/register?Phone=xxx         
//${Base_Url}/register?Phone=           队友
//http://redrock.udday.cn:2022/captcha/sent?phone=13xxx  网易云

async function getMessage(phone) {
    const require = await fetch(`${Base_Url}register?Phone=${phone}`)

    return await require.json()
}


async function getToken(url, token) {
    const require = await fetch(`${Base_Url}${url}`, {
        // credentials: 'include',
        headers: {
            'Authorization': `${token}`
        }
    })
    return await require.json()
}


async function out(url, token) {
    let res = await fetch(`${Base_Url}${url}`, {
        headers: {
            'Authorization': `${token}`
        }
    })
    return await res.json()
}