requirejs(["ciphers/caesar"]);

const showResult = (data, handler) => {
  handler(data);
};
window.onload = () => {
  document.getElementById(
    "caesar-cipher-encrypt-button"
  ).onclick = async () => {
    const response = await caesarCipherEncrypt();
    showResult(response.data, caesarCipherEncryptHandler);
  };
};
