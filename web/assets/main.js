window.onload = init;

function init() {
    checkAuth()
    addEventListeners()
    ListenURLVisitEvents()
}

function checkAuth() {
    if (location.pathname == "/login" || location.pathname == "/register") {
        return
    }
    
    let token = localStorage.getItem("accessToken")
    if (!token) {
        location.pathname = "/login"
        return
    }

    let user = JSON.parse(localStorage.getItem("user"))

    fetch(`/api/v1/users/${user.ID}`, {
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        }
    })
    .then(res => res.json())
    .then(data => {
        localStorage.setItem("user", JSON.stringify(data))
        document.getElementById("username").innerHTML = `@${data.name}`
    })
    .catch(err =>  {
        console.log(err)
        location.pathname = "/login"
    })
}

function addEventListeners() {
    let createURLForm = document.getElementById("js_url-form")
    if (createURLForm) {
        createURLForm.addEventListener("submit", handleGenerateURL, false)
    }

    let loginForm = document.getElementById("js_login-form")
    if (loginForm) {
        loginForm.addEventListener("submit", handleLogin, false)
    }

    let registerForm = document.getElementById("js_register-form")
    if (registerForm) {
        registerForm.addEventListener("submit", handleRegister, false)
    }

    let logoutBtn = document.getElementById("logout-btn")
    if (logoutBtn) {
        logoutBtn.addEventListener("click", handleLogout, false)
    }
}

function handleGenerateURL(e) {
    e.preventDefault()

    let originalURL = document.getElementById("url").value
    let customKey = document.getElementById("custom_url").value

    if (originalURL.length === 0 || !isValidHttpUrl(originalURL)) {
        alert("Please enter a valid URL")
        return
    }

    fetch("/api/v1/urls", {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            "Content-Type": "application/json",
            "Authorization": `Bearer ${getAccessToken()}`
        },
        body: JSON.stringify({
            original_url: originalURL,
            custom_key: customKey
        })
    })
    .then(res => res.json())
    .then(data => handleResponse(data))
    .catch(err => handleError(err))

    function handleResponse(data) {
        location.reload()
    }

    function handleError(err) {
        console.log(err)
        alert(err.message)
    }
}

function getAccessToken() {
    return localStorage.getItem("accessToken")
}

function isValidHttpUrl(string) {
    let url;
    try {
      url = new URL(string);
    } catch (_) {
      return false;
    }
    return url.protocol === "http:" || url.protocol === "https:";
}

function FormatDate(date) {
    options = {
        year: "numeric",
        month: "numeric",
        day: "numeric",
        hour: "numeric",
        minute: "numeric",
        second: "numeric",
        hour12: false,
        timeZone: "Europe/Moscow",
      };

    return new Intl.DateTimeFormat('en-GB', options).format(date);
}

function handleRegister(e) {
    e.preventDefault()

    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let password = document.getElementById("password").value

    fetch("/api/v1/users", {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            name,
            email,
            password
        })
    })
    .then(res => res.json())
    .then(data => {
        Login(email, password)
    })
    .catch(err => {
        alert(err.message)
    })
}

function handleLogin(e) {
    e.preventDefault()

    let login = document.getElementById("login").value
    let password = document.getElementById("password").value

    Login(login, password)
}

function Login(login, password) {
    fetch("/api/v1/auth/login", {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            login,
            password
        })
    })
    .then(res => res.json())
    .then(data => {
        localStorage.setItem("accessToken", data.token)
        localStorage.setItem("user", JSON.stringify(data.user))
        location.pathname = "/"
    })
    .catch(err => {
        alert(err.message)
    })
}

function handleLogout(e) {
    e.preventDefault()

    localStorage.removeItem("accessToken")
    localStorage.removeItem("user")
    location.pathname = "/login"
}

function ListenURLVisitEvents() {
    let table = document.getElementById("visits-table");
    if (!table) {
        return 
    }

    let urlId = table.dataset.urlId
    let eventSource = new EventSource(`/api/v1/urls/${urlId}/events`);

    eventSource.onmessage = (event) => {
        console.log(event)
        console.log(event.data);
        // document.getElementById("random-number").innerHTML = event.data;
    }
}