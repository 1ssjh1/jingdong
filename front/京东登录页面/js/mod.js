 const Base_Url = "https://sanser.ltd"
 export {get, post, postForm, getMessage, postFormDate, getToken }

 // -------------------------------------
 // -------------------------------------
 async function post(url, data) {
     const require = await fetch(`${Base_Url}${url}`, {
         method: "POST",
         credentials: "include",
         headers: {
             'Content-Type': `application/json`
         },
         body: JSON.stringify(data)
     })

     return await require.json()
 }

 async function postForm(url, data) {
     const require = await fetch(`${Base_Url}${url}`, {
         method: "POST",
         //  credentials: "include",
         headers: {
             'Content-Type': `application/json`,
             //  "Accept - Encoding": "none"
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


 async function getToken(url, Token) {
     const require = await fetch(`${Base_Url}${url}`, {
         headers: {
             "Authorization": "Bearer" + localStorage.getItem("token")
         }
     })
     return await require.json()
 }