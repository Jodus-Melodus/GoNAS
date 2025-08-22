
document.addEventListener("DOMContentLoaded", () => {
    const logoutLink = document.getElementById("logout-link");
    const logoutForm = document.getElementById("logout-form");


    if (logoutLink && logoutForm) {
        logoutLink.addEventListener("click", function (e) {
            e.preventDefault();
            logoutForm.submit();
        });
    }
});

const fileInput = document.getElementById('file');
const fileSelected = document.getElementById('file-selected');

if ((fileInput != null) && (fileSelected != null)) {


    fileInput.addEventListener('change', () => {
        if (fileInput.files.length > 0) {
            fileSelected.textContent = fileInput.files[0].name;
        } else {
            fileSelected.textContent = 'No file chosen';
        }
    });
}