const messagesContainer = document.getElementById("messages");
const inputField = document.getElementById("input-field");
const recordButton = document.getElementById("record-button");
let isRecording = false;

recordButton.addEventListener("click", function() {
    if (isRecording) {
        // 结束录音
        isRecording = false;
        this.innerHTML = "录音";
    } else {
        // 开始录音
        isRecording = true;
        this.innerHTML = "结束录音";
    }
});

inputField.addEventListener("keyup", function(event) {
    if (event.key === "Enter") {
        const message = this.value;
        messagesContainer.innerHTML += `<div>${message}</div>`;
        this.value = "";
        // 调用 func2 或 func3
    }
});
