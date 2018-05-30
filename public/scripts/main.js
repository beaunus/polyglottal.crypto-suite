requirejs(["ciphers/caesar", "ciphers/scytale"]);

const showResult = (data, handler) => {
  handler(data);
};
window.onload = () => {
  document.getElementById("caesar-encrypt-button").onclick = async () => {
    const response = await caesarEncrypt();
    showResult(response.data, caesarEncryptHandler);
  };
  document.getElementById("scytale-encrypt-button").onclick = async () => {
    const response = await scytaleEncrypt();
    showResult(response.data, scytaleEncryptHandler);
  };
};
