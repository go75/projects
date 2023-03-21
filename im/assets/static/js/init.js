var xhr = new XMLHttpRequest()
function doReq(method, url, callback) {
    xhr.open(method, url, true)
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send();
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            callback(xhr.responseText)
        }
    }
}
window.onload = function() {
    let btn = document.getElementsByTagName('button').item(0)
    let a = document.getElementsByTagName('a').item(0)
    let title = document.getElementsByTagName('h1').item(0)
    let info = document.getElementById('user-info')
    let check = document.createElement('input')
    let br = document.createElement('br')
    let form = document.getElementById('user-info')
    check.id = 'check'
    check.type = 'password'
    check.name = 'check'
    check.placeholder = '确认密码'
    a.onclick = function() {
        if (title.innerText == '登录') {
            info.removeChild(btn)
            title.innerText = '注册'
            btn.innerText = '注册'
            a.innerText = '已有帐号?点击登录'
            form.action = 'http://127.0.0.1:9999/user/regist'
            info.appendChild(check)
            info.appendChild(br)
            info.appendChild(btn)
        } else {
            info.removeChild(br)
            info.removeChild(check)
            title.innerText = '登录'
            btn.innerText = '登录'
            a.innerText = '没有帐号?点击注册'
            form.action = 'http://127.0.0.1:9999/user/init'
        }
    }
    btn.onclick = function() {
        doReq('get', 'http://127.0.0.1:9999/user/login', function(text) {
            localStorage.setItem('im-token', text)
            location.replace('http://127.0.0.1:9999')
        })
    }
}