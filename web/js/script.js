// TODO: rewrire to jQuery
// TODO: rename file
const applicantForm = document.getElementById('image-form');
const imagePreview = document.getElementById('image-preview');
const imageInput = document.querySelector('input[name="image"]');
const btnThemeSwitch = document.querySelector('.btn-theme-switch')
const btnSubmit = document.querySelector('#image-form .btn');
const btnSubmitSpinner = document.querySelector('#image-form .btn .spinner');
const toastContainer = document.querySelector('.toast-container');

function showMessage(title, message) {

	const messegeHTML = `
	<div class="toast show" role="alert" aria-live="assertive" aria-atomic="true">
		<div class="toast-header">
			<strong class="me-auto">${title}</strong>
			<button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
		</div>
		<div class="toast-body">
			${message}
		</div>
	</div>`

	toastContainer.insertAdjacentHTML('beforeend', messegeHTML);
}

function handleFormSubmit(event) {
	event.preventDefault();
	btnSubmit.disabled = true;
	btnSubmitSpinner.classList.remove("visually-hidden");
	var data = new FormData(applicantForm);

	fetch('/image', {
		method: 'POST',
		body: data
	})
		.then(res => {
			if (!res.ok) {
				return res.text()
			}
			return res.blob();
		})
		.then(blob => {
			if (typeof blob == "string") {
				showMessage("Server", blob);
			} else if (typeof blob == "object") {
				const imgUrl = URL.createObjectURL(blob);
				imagePreview.src = imgUrl;
			}
			btnSubmit.disabled = false;
			btnSubmitSpinner.classList.add("visually-hidden")
		})
}

function checkFileType(file) {
	const allowedTypes = ['image/png', 'image/jpeg'];
	if (allowedTypes.includes(file.type)) {
		return true;
	} else {
		return false;
	}
}

function changeImagePreview() {
	const file = this.files[0];
	if (checkFileType(file)) {
		const reader = new FileReader();
		reader.addEventListener("load", function () {
			imagePreview.src = reader.result;
		});

		if (file) {
			reader.readAsDataURL(file);
		}
	} else {
		showMessage("ImageEditor", "Неверный тип файла. Пожалуйста, выберите файл в формате PNG или JPEG.");
		throw new Error('Network response was not ok');
	}
}
function themeSwitch() {
	let newTheme = null;
	let oldTheme = null;
	if (document.documentElement.getAttribute('data-bs-theme') === 'dark') {
		newTheme = 'light';
		oldTheme = 'dark';
	}
	else {
		newTheme = 'dark';
		oldTheme = 'light';
	}

	const oldClass = 'border-' + oldTheme + '-subtle';
	const elements = document.querySelectorAll('.' + oldClass);
	console.log(elements);
	for (const object of elements) {
		console.log(object);
		object.classList.remove(oldClass);
		object.classList.add('border-' + newTheme + '-subtle');
	}
	document.documentElement.setAttribute('data-bs-theme', newTheme);
}

btnThemeSwitch.addEventListener('click', themeSwitch);
applicantForm.addEventListener('submit', handleFormSubmit);
imageInput.addEventListener('change', changeImagePreview);
