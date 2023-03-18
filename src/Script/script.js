function openPopUp(element) {
    element.nextElementSibling.style.display = "block";
    document.body.style.overflow = "hidden";
}

function closePopUp(element) {
    element.parentElement.parentElement.style.display = "none";
    document.body.style.overflow = "auto";
}