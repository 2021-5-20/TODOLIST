document.getElementById("login-form").addEventListener("submit", async (e) => {
  e.preventDefault();

  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;
  const resultElement = document.getElementById("result");

  resultElement.textContent = "正在登录，请稍候...";

  try {
    const res = await fetch("/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password })
    });

    const result = await res.json();

    if (res.ok) {
      localStorage.setItem("token", result.token);
      resultElement.textContent = "登录成功，跳转中...";
      setTimeout(() => {
  window.location.href = "/memo";
}, 500);
    } else {
      resultElement.textContent = "登录失败：" + (result.error || "用户名或密码错误");
    }
  } catch (err) {
    resultElement.textContent = "请求失败：" + err.message;
  }
});