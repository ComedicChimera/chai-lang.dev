class FeatherIcon extends HTMLElement {
    constructor() {
        super()
    }

    connectedCallback() {
        let iconName = this.getAttribute("name")
        let dim = this.hasAttribute("dim") ? this.getAttribute("dim") : "50px"

        this.innerHTML = `<div class="feather-icon">${feather.icons[iconName].toSvg({
            stroke: "#212121",
            "stroke-width": "1.5px",
            height: dim,
            width: dim,
        })}</div>`
    }
}

customElements.define('feather-icon', FeatherIcon)