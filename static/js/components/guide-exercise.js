class GuideExercise extends HTMLElement {
    constructor() {
        super()
    }

    connectedCallback() {
        if (this.innerText === null) return

        let jdata = JSON.parse(this.innerText)

        if (jdata.solution.type == "program") {
            fetch("/api/get-solution-code/" + jdata.solution.url).then(resp => {
                if (resp.status == 200) {
                    resp.text().then(html => this.render(jdata, 
                        `<pre x-init="$nextTick(() => Prism.highlightElement($el))"
                        class="language-chai"><code>${html}</code></pre>`
                    ))
                }
            })
        } else {
            this.render(jdata, jdata.solution.text)
        }   
    }

    render(jdata, solutionHTML) {
        this.innerHTML = `
            <div class="guide-exercise" x-data="{ hint_open: false, solution_open: false }">
                <div class="exercise-top-bar">
                    <div class="exercise-title">Exercise ${jdata.label}</div>
                    <div class="exercise-buttons">
                        <feather-icon id="hint-button" name="help-circle" dim="1.5rem" @click="hint_open = !hint_open"></feather-icon>
                        <feather-icon id="solution-button" name="zap" dim="1.5rem" @click="solution_open = !solution_open"></feather-icon>
                    </div>                            
                </div>
                <div class="exercise-content">
                    ${jdata.content}
                </div>
                <div class="exercise-hint" x-show="hint_open" x-transition.duration.500ms>
                    <div class="hint-title">Hint</div>
                    <div class="hint-text">${jdata.hint}</div>
                </div>
                <div class="exercise-solution" x-show="solution_open" x-transition.duration.500ms>
                    <div class="solution-title">Solution</div>
                    <div class="solution-content">
                        ${solutionHTML}
                    </div>
                </div>
            </div>
        `
    }
}

customElements.define('guide-exercise', GuideExercise);