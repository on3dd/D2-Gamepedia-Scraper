window.onload = function () {
    let playBtns = document.getElementsByClassName("audiobtn");
    let audioSrcs = document.getElementsByTagName("audio");

    for (let i = 0; i < playBtns.length; i++) {
        playBtns[i].addEventListener("click", () => {
            let audio = audioSrcs[i];
            if (audio.paused || audio.ended) audio.play();
            else audio.pause();
        });
    }
}