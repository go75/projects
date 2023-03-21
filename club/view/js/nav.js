window.onload = function () {
    var index = 0;
    var nav = document.getElementById('nav');
    var list = ['综合', 'all', '前端', 'front-end', '后端', 'back-end', '云原生', 'cloud-native', '大数据', 'big-data', '区块链', 'block-chain', '游戏开发', 'game', '人工智能', 'ai', '杂谈', 'other'];
    var _loop_1 = function (i) {
        var a = document.createElement('a');
        a.innerText = list[i];
        a.id = list[i + 1];
        a.style.cursor = 'pointer';
        a.onclick = function () {
            document.getElementById(list[index + 1]).style.color = 'black';
            index = i;
            a.style.color = 'cornflowerblue';
        };
        a.onmouseover = function () {
            a.style.color = 'green';
        };
        a.onmouseout = function () {
            console.log(i, index);
            if (i == index) {
                a.style.color = 'cornflowerblue';
            }
            else {
                a.style.color = 'black';
            }
        };
        if (i == 0) {
            a.style.color = 'cornflowerblue';
        }
        var li_1 = document.createElement('li');
        li_1.appendChild(a);
        nav.appendChild(li_1);
    };
    for (var i = 0; i < list.length; i += 2) {
        _loop_1(i);
    }
    // <li>
    //      <input type="text" id="keyInput" placeholder="   输入关键字" />
    //      <svg width="45px" height="45px" viewBox="0 0 17 17" fill="none" xmlns="http://www.w3.org/2000/svg" id="searchBtn" onclick="searchFunc()">
    //            <path fill-rule="evenodd" clip-rule="evenodd" d="M16.3451 15.2003C16.6377 15.4915 16.4752 15.772 16.1934 16.0632C16.15 16.1279 16.0958 16.1818 16.0525 16.2249C15.7707 16.473 15.4456 16.624 15.1854 16.3652L11.6848 12.8815C10.4709 13.8198 8.97529 14.3267 7.44714 14.3267C3.62134 14.3267 0.5 11.2314 0.5 7.41337C0.5 3.60616 3.6105 0.5 7.44714 0.5C11.2729 0.5 14.3943 3.59538 14.3943 7.41337C14.3943 8.98802 13.8524 10.5087 12.8661 11.7383L16.3451 15.2003ZM2.13647 7.4026C2.13647 10.3146 4.52083 12.6766 7.43624 12.6766C10.3517 12.6766 12.736 10.3146 12.736 7.4026C12.736 4.49058 10.3517 2.1286 7.43624 2.1286C4.50999 2.1286 2.13647 4.50136 2.13647 7.4026Z" fill="currentColor"></path>
    //      </svg>
    // </li>
    var li = document.createElement('li');
    li.innerHTML = '<input type="text" id="keyInput" placeholder="   输入关键字" /><svg width="45px" height="45px" viewBox="0 0 17 17" fill="none" xmlns="http://www.w3.org/2000/svg" id="searchBtn" onclick="searchFunc()"><path fill-rule="evenodd" clip-rule="evenodd" d="M16.3451 15.2003C16.6377 15.4915 16.4752 15.772 16.1934 16.0632C16.15 16.1279 16.0958 16.1818 16.0525 16.2249C15.7707 16.473 15.4456 16.624 15.1854 16.3652L11.6848 12.8815C10.4709 13.8198 8.97529 14.3267 7.44714 14.3267C3.62134 14.3267 0.5 11.2314 0.5 7.41337C0.5 3.60616 3.6105 0.5 7.44714 0.5C11.2729 0.5 14.3943 3.59538 14.3943 7.41337C14.3943 8.98802 13.8524 10.5087 12.8661 11.7383L16.3451 15.2003ZM2.13647 7.4026C2.13647 10.3146 4.52083 12.6766 7.43624 12.6766C10.3517 12.6766 12.736 10.3146 12.736 7.4026C12.736 4.49058 10.3517 2.1286 7.43624 2.1286C4.50999 2.1286 2.13647 4.50136 2.13647 7.4026Z" fill="currentColor"></path></svg>';
    nav.appendChild(li);
    li = document.createElement('li');
    var img = document.createElement('img');
    img.className = 'header-img';
    img.src = '/view/img/login.png';
    li.appendChild(img);
    nav.appendChild(li);
    li = document.createElement('li');
    li.style.color = 'cornflowerblue';
    li.innerText = '咸鱼的自我修养';
    nav.appendChild(li);
    var input = document.getElementById('keyInput');
    var btn = document.getElementById('searchBtn');
    input.onmouseover = function () {
        input.style.filter = 'brightness(100%)';
    };
    input.onmouseout = function () {
        input.style.filter = 'brightness(90%)';
    };
    btn.onmouseover = function () {
        btn.style.color = 'cornflowerblue';
    };
    btn.onmouseout = function () {
        btn.style.color = 'black';
    };
};
