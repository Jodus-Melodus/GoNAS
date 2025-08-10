document.querySelectorAll(".delete-button").forEach(button => {
    button.addEventListener("click", () => {
        const name = button.dataset.name;

        fetch("/delete", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ name })
        })
            .then(res => {
                if (res.ok) {
                    alert("Delete successful!");
                    location.reload();
                } else {
                    alert("Delete failed.");
                }
            })
            .catch(err => console.error(err));
    })
});

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