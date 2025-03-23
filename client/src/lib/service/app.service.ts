import axios from "axios";

async function checkServerConnection() {
  try {
    await axios.get("http://localhost:3000/api/health");
    return true;
  } catch (error) {
    console.error("Sunucu bağlantısı hatalı:", error);
    return false;
  }
}

export { checkServerConnection };
