// Main entry point
document.addEventListener('DOMContentLoaded', start);

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
	var loginForm = document.getElementById('form');
	loginForm.addEventListener('submit', function(event) {
		event.preventDefault();
		userLogin();
	});
}

function userLogin() {
	var userName = document.querySelector('input[name="userName"]');
	var userPassWord = document.querySelector('input[name="userPassword"]');
	var nameVal = userName.value,
		passwordVal = userPassWord.value,
		xsessionVal = ".xinitrc";
	// Wails v2: bound struct methods are available at window.go.main.StructName.MethodName
	window.go.main.LoginHandler.Login(nameVal, passwordVal, xsessionVal).then(function() {
		console.log("login submitted");
	}).catch(function(err) {
		console.error("login error:", err);
		loginCheck(false);
	});
}

function loginCheck(isLogin) {
	var modal = document.querySelector('.container-msg-modal');
	var errorMsg = document.querySelector('.container-modal-content--error');
	var successMsg = document.querySelector('.container-modal-content--success');
	var loginForm = document.getElementById('form');
	var modalContent = document.querySelectorAll('.container-modal-content');
	modal.classList.add('enabled');
	if(isLogin) {
		successMsg.classList.add('enabled');
	} else {
		errorMsg.classList.add('enabled');
	}

	setTimeout(function() {
		modal.classList.remove('enabled');
		loginForm.reset();
		modalContent.forEach(function(content) {
			content.classList.remove('enabled');
		});
	}, 3000);
}
