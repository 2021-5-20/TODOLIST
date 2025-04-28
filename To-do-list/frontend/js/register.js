document.getElementById("register-form").addEventListener("submit", async (e) => {
    e.preventDefault();
  
    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();
    const resultElement = document.getElementById("result");
  
    // 👉 前端简单校验
    if (!username || !password) {
      resultElement.textContent = "用户名和密码不能为空！";
      return;
    }
    if (password.length < 6) {
      resultElement.textContent = "密码至少需要6位！";
      return;
    }
  
    resultElement.textContent = "正在注册，请稍后...";
  
    try {
      const res = await fetch("/api/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
      });
  
      const result = await res.json();
  
      if (res.ok) {
        resultElement.classList.remove("text-danger");
        resultElement.classList.add("text-success");
        resultElement.textContent = "注册成功，1秒后跳转登录...";
        setTimeout(() => {
          window.location.href = "/";
        }, 1000);
      } else {
        resultElement.textContent = "注册失败：" + (result.error || "未知错误");
      }
    } catch (err) {
      resultElement.textContent = "请求失败：" + err.message;
    }
  });