requirejs(["ciphers/caesar"]);

const showResult = (data, handler) => {
  handler(data);
};
window.onload = () => {
  document.getElementById("caesar-encrypt-button").onclick = async () => {
    const response = await caesarEncrypt();
    showResult(response.data, caesarEncryptHandler);
  };
};
