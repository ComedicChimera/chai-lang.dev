class GuideExercise extends HTMLElement {
    constructor() {
        super()
    }

    connectedCallback() {
        let exLabel = this.getAttribute('label')
        if (exLabel === null) {
            return;
        }
        
        let fetchProm = fetch("/api/get-guide-exercise?label=" + exLabel)
        fetchProm.then(resp => {
            if (resp.status != 200) {
                this.throwError(exLabel, resp.statusText)
                return
            }

            resp.json().then(jdata => {
                this.innerHTML = `
                    <div class="guide-exercise" x-data="{ hint_open: false, solution_open: false }">
                        <div class="exercise-top-bar">
                            <div class="exercise-title">Exercise ${exLabel}</div>
                            <div class="exercise-buttons">
                                <feather-icon id="hint-button" name="help-circle" dim="1.5rem" @click="hint_open = !hint_open"></feather-icon>
                                <feather-icon id="solution-button" name="zap" dim="1.5rem"></feather-icon>
                            </div>                            
                        </div>
                        <div class="exercise-content">
                            ${jdata.content}
                        </div>
                        <div class="exercise-hint" x-show="hint_open" x-transition.duration.500ms>
                            <div class="hint-title">Hint</div>
                            <div class="hint-text">${jdata.hint}</div>
                        </div>
                    </div>
                `
            }).catch(e => this.throwError(exLabel, e))  
        }).catch(e => this.throwError(exLabel, e))
    }

    throwError(exLabel, e) {
        console.log("Failed to retrieve exercise " + exLabel)
        console.log(e)
    }
}

customElements.define('guide-exercise', GuideExercise);