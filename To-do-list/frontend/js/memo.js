async function loadMemos() {
  const token = localStorage.getItem("token");
  const container = document.getElementById("memo-list");

  if (!token) {
    alert("未登录或登录已失效，请重新登录");
    window.location.href = "/";
    return;
  }

  // 👉 加载时显示 Bootstrap 小转圈动画
  container.innerHTML = `
    <div class="d-flex justify-content-center mt-5">
      <div class="spinner-border text-primary" role="status" style="width: 3rem; height: 3rem;">
        <span class="visually-hidden">Loading...</span>
      </div>
    </div>
  `;

  try {
    const res = await fetch("/api/memo", {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${token}`
      }
    });

    if (!res.ok) {
      alert("获取失败，可能登录失效。");
      window.location.href = "/";
      return;
    }

    const memos = await res.json();
    renderMemos(memos);
  } catch (err) {
    alert("请求异常：" + err.message);
  }
}
  
  function renderMemos(memos) {
    const container = document.getElementById("memo-list");
    container.innerHTML = "";
  
    if (memos.length === 0) {
      container.innerHTML = "<p>暂无备忘录，点击右上角添加吧！</p>";
      return;
    }
  
    memos.forEach((memo, index) => {
      const createdAt = formatDate(memo.created_at);
      const div = document.createElement("div");
      div.className = "col-md-4";
  
      div.innerHTML = `
        <div class="card h-100 shadow-sm">
          <div class="card-body d-flex flex-column">
            <h5 class="card-title">#${index + 1}</h5>
            <p class="card-text flex-grow-1">${memo.content}</p>
            <p class="text-muted mb-2" style="font-size: 0.9em;">创建时间：${createdAt}</p>
            <button class="btn btn-sm btn-danger mt-auto" onclick="deleteMemo(${memo.id})">删除</button>
          </div>
        </div>
      `;
  
      container.appendChild(div);
    });
  }
  function addMemo() {
    const modal = new bootstrap.Modal(document.getElementById('addMemoModal'));
    document.getElementById("memoContent").value = ""; // 清空输入框
    modal.show();
  }
  async function submitMemo() {
    const content = document.getElementById("memoContent").value.trim();
    if (!content) {
      alert("请输入内容！");
      return;
    }
  
    const token = localStorage.getItem("token");
  
    const res = await fetch("/api/memo", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`
      },
      body: JSON.stringify({ content })
    });
  
    if (res.ok) {
      const modal = bootstrap.Modal.getInstance(document.getElementById('addMemoModal'));
      modal.hide();  // 关闭弹窗
      alert("添加成功！");
      loadMemos();   // 重新加载列表
    } else {
      const err = await res.json();
      alert("添加失败：" + err.error);
    }
  }
  async function deleteMemo(id) {
    // 用浏览器自带的确认弹窗确认一下
    if (!confirm("确定要删除这条备忘录吗？删除后无法恢复！")) return;
  
    const token = localStorage.getItem("token");
  
    const res = await fetch(`/api/memo/${id}`, {
      method: "DELETE",
      headers: {
        "Authorization": `Bearer ${token}`
      }
    });
  
    if (res.ok) {
      alert("删除成功！");
      loadMemos(); // 刷新列表
    } else {
      const err = await res.json();
      alert("删除失败：" + (err.error || "未知错误"));
    }
  }

function logout() {
    localStorage.removeItem("token");
    window.location.href = "/";
  }

  function formatDate(dateString) {
    const date = new Date(dateString);
    const year = date.getFullYear();
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}`;
  }
  window.onload = loadMemos;