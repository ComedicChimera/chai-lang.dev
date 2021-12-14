// add syntax highlighting
for (let elem of document.getElementsByTagName("pre")) {
    if (elem.classList.length == 0) {
        elem.classList.add("language-chai")
    } else {
        elem.setAttribute('class', `language-${elem.getAttribute('class')}`)
    }
}

// add subheading references
for (let i in document.getElementsByTagName("h2")) {
    let item = sections[i]
    item.innerHTML = `<a name="section${i}"/> ${item.innerText}`
}