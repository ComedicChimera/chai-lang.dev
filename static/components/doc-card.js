class DocCard extends HTMLElement {
    constructor() {
        super()
    }

    connectedCallback() {
        let docName = this.getAttribute('name')
        let docHref = this.getAttribute('href')

        this.innerHTML = `
            <div class="doc-card">
                <div class="doc-name"><a href="${docHref}">${docName}</a></div>
                <div class="doc-text">${this.innerText}</div>
            </div>
        `
    }
}

customElements.define('doc-card', DocCard);