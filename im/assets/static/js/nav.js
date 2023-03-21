var xhr = new XMLHttpRequest()
var content = document.getElementById('')

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

function friendFunc() {
    doReq('get', 'http://127.0.0.1:9999/user/friend', function(text) {
        console.log(text)
    })
}

function groupFunc() {
    doReq('get', 'http://127.0.0.1:9999/group', function(text) {
        console.log(text)
    })
}

function addFunc() {
    clearContent()
    let user = document.createElement('option')
    user.value = 'user'
    user.innerText = '查找用户'
    let group = document.createElement('option')
    group.value = 'group'
    group.innerHTML = '查找群聊'
    let type = document.createElement('select')
    type.appendChild(user)
    type.appendChild(group)
    content.appendChild(type)

    let input = document.createElement('input')
    content.appendChild(input)

    let btn = content.createElement('button')
    btn.onclick = function() {
        doReq('get', 'http://127.0.0.1:9999/'+type.options[type.selectedIndex].value+'/add/'+input.value, function(text){
            console.log(text)

        })
    }
    content.appendChild(btn)
}

function clearContent() {
    while (content.childNodes.length > 0) {
        content.removeChild(content.childNodes.item(0))
    }
}

window.onload = function() {
    friendFunc()
    let friend = document.getElementById('friend')
    friend.onclick = friendFunc
    let group = document.getElementById('group')
    let add = document.getElementById('add')
    content = document.getElementById('content')
    add.onclick = addFunc()
    // ws = new WebSocket("ws://127.0.0.1:9999/upgrade")
    // ws.onopen = function () {
    //     console.log('连接建立')
    // }
    // ws.onclose = function () {
    //     alert('连接断开')
    //     ws = null;
    // }
    // ws.onmessage = function (evt) {
    //     evt.data
    // }
    // ws.onerror = function (evt) {
    //     alert("ERROR: " + evt.data);
    // }
}