

document.querySelectorAll(".download-button").forEach(button => {
    button.addEventListener("click", () => {
        const name = button.dataset.name;

        // fetch("/download-file", {
        //     method: "POST",
        //     headers: {
        //         "Content-Type": "application/json"
        //     },
        //     body: JSON.stringify({ name })
        // })
        //     .then(res => {
        //         if (res.ok) {
        //             alert("Download successful!");
        //             // reaload
        //         } else {
        //             alert("Download failed.");
        //         }
        //     })
        //     .catch(err => console.error(err));
    })
});

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