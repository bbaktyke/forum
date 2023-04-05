var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

let addcomment = document.querySelector(".add__comment")
let comment=document.querySelector(".comment")
addcomment.addEventListener("click", ()=> {
	comment.classList.toggle("active")
})