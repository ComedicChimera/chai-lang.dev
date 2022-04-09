class FeatherIcon extends HTMLElement {
    constructor() {
        super()
    }

    connectedCallback() {
        this.render()
    }

    attributeChangedCallback() {
        this.render()
    }

    static get observedAttributes() {
        return ["name"]
    }

    render() {
        if (this.hasAttribute("name")) {
            let iconName = this.getAttribute("name")
            let dim = this.hasAttribute("dim") ? this.getAttribute("dim") : "50px"
            let color = this.hasAttribute("color") ? this.getAttribute("color") : "#212121"

            this.innerHTML = `<div class="feather-icon">${feather.icons[iconName].toSvg({
                stroke: color,
                "stroke-width": "1.5px",
                height: dim,
                width: dim,
            })}</div>`
        } 
    }
}

customElements.define('feather-icon', FeatherIcon)