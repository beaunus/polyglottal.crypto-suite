const caesarEncryptHandler = data => {
  const element = document
    .getElementById("caesar")
    .getElementsByClassName("cipher")[0]
    .getElementsByClassName("row")[0]
    .getElementsByClassName("result")[0];
  element.innerText = data.Ciphertext;
};

const caesarEncrypt = async () => {
  return await axios.get("/api/v1/caesar", {
    params: {
      plaintext: document.getElementById("caesar-encrypt-plaintext").value,
      shift: document.getElementById("caesar-encrypt-shift").value
    }
  });
};

document.getElementById("caesar-encrypt-button").onclick = async () => {
  const response = await caesarEncrypt();
  showResult(response.data, caesarEncryptHandler);
};
