class SectionTitle extends HTMLElement {
    constructor() {
        super()
    }   

    connectedCallback() {
        let titleString = this.innerText

        this.innerHTML = `
            <div class="title-container">
                <span class="title-text">${titleString}</span>
                <svg class="title-svg" width="80px" height="5px">
                    <rect class="title-svg-left" width="60px" height="5px"/>
                    <rect class="title-svg-right" x="60px" width="20px" height="5px"/>
                </svg>
            </div>
        `
    }
}

customElements.define('section-title', SectionTitle)