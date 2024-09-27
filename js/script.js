// On load the loader will disappear and the content will fade in

window.onload = function() {
    const loader = document.getElementById('loader');
    const content = document.body;

    loader.style.display = 'none';

    content.classList.add('fade-in');
};