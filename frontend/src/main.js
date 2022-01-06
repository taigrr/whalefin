import 'core-js/stable';
const runtime = require('@wailsapp/runtime');
// Main entry point
function start() {
	// Ensure the default app div is 100% wide/high
	var app = document.getElementById('app');
	app.style.width = '100%';
	app.style.height = '100%';

	// Inject html
	app.innerHTML = `
	<div class='logo'></div>
	  <div class="container">
		<div class="container-msg-modal">
		  <div class="container-modal-content--success container-modal-content">
			<span id="hostname">Welcome!</span>
		  </div>
		<div class="container-modal-content--error container-modal-content"><span>Failed Login</span></div>
	  </div>
	  <form action="" id="form">
		<div class="container-form-userName container-form-input">
		  <label for="userName">User name</label>
		  <input type="text" placeholder="User Name" name="userName" required>
		</div>
		<div class="container-form-userPassword container-form-input">
		  <label for="userPasswrod">Password</label>
		  <input type="password" placeholder="User Password" name="userPassword" onfocus="this.value=''" required>
		</div>
		<button type="submit" class="js-form-btn">Submit</button>
	  </form>
	</div>
	`;
	var errorMsg = document.querySelector('.container-modal-content--error')
	var successMsg = document.querySelector('.container-modal-content--success')
	var loginForm = document.getElementById('form')
	var modal = document.querySelector('.container-msg-modal')
	var modalContent = document.querySelectorAll('.container-modal-content')
	loginForm.addEventListener('submit', function(event) {
		event.preventDefault()
		userLogin()
	})
};
const myLogin = {
	userName: 'codepen',
	password: 'codepen'
}

function userLogin() {
	var userName = document.querySelector('input[name="userName"]')
	var userPassWord = document.querySelector('input[name="userPassword"]')
	var nameVal = userName.value,
		passwordVal = userPassWord.value,
		xsessionVal = ".xinitrc"
	window.backend.login({username: nameVal, password:passwordVal, xsession: xsessionVal}).then(result => {
		console.log(result);
	});
}

function loginCheck(isLogin) {
	modal.classList.add('enabled')  
	if(isLogin) {
		successMsg.classList.add('enabled')
	} else {
		errorMsg.classList.add('enabled')
	}

	setTimeout(function() {
		modal.classList.remove('enabled')
		loginForm.reset();
		modalContent.forEach(function(content) {
			content.classList.remove('enabled')
		});
	}, 3000)
}
// We provide our entrypoint as a callback for runtime.Init
runtime.Init(start);
