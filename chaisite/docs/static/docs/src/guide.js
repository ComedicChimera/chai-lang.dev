import GuideAside from './guide-src/GuideAside.svelte'
import GuideSummary from './guide-src/GuideSummary.svelte'
import ChaiTitle from 'common/ChaiTitle.svelte'

import prism from 'prismjs'

import './guide-src/guide-styles.scss'

new GuideAside({
    target: document.getElementById("guide-aside")
})

prism.languages.chai = window.language_chai

for (let parentElement of document.getElementsByTagName('pre')) {
    let element = parentElement.childNodes[0]

    element.classList.add("language-chai")
    prism.highlightElement(element, false)
}

let articleElement = document.getElementsByTagName("article")[0]

for (let h1Element of articleElement.getElementsByTagName("h1")) {
    new ChaiTitle({
        target: h1Element,
        props: {
            titleString: h1Element.innerHTML
        },
        hydrate: true
    })
}

new GuideSummary({
    target: document.getElementById("guide-summary")
})

for (let h2 of articleElement.getElementsByTagName("h2")) {
    h2.innerHTML = `<a name="${h2.innerHTML.replaceAll(" ", "-").toLowerCase()}">${h2.innerHTML}</a>`
}

for (let table of articleElement.getElementsByTagName("table")) {
    table.setAttribute("border", "1")
    table.setAttribute("cellspacing", "0px")
}
