const caesarCipherEncryptHandler = data => {
  const element = document
    .getElementById("caesar-cipher")
    .getElementsByClassName("encrypt")[0]
    .getElementsByClassName("result")[0];
  element.innerText = data.Ciphertext;
};

const caesarCipherEncrypt = async () => {
  return await axios.get("/api/v1/caesarCipher/encrypt", {
    params: {
      plaintext: document.getElementById("caesar-cipher-encrypt-plaintext")
        .value,
      shift: document.getElementById("caesar-cipher-encrypt-shift").value
    }
  });
};
