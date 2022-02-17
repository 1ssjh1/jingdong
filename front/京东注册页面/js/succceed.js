let username = document.querySelector("i")
let username_true = JSON.parse(localStorage.getItem("username"))

let button = document.querySelector("button")


username.innerHTML = username_true
localStorage.removeItem("username")





button.addEventListener("click", (e) => {
    e.preventDefault() //组织刷新
    console.log(1);

    // http://127.0.0.1:5501/index.html   登录页面
    // http://127.0.0.1:5502/html/index.html    首页
    // ../京东登录页面/index.html
    window.location = "http://127.0.0.1:5502/html/index.html"
        // location.replace(location.href);

})