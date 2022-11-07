const formEl = document.getElementById("form");
const resultEl = document.getElementById("result");
const resultGroupEl = document.getElementById("result-group");
const loaderEl = document.getElementById("loader");
const aEl = document.getElementById("a");
const bEl = document.getElementById("b");
const mathExpressionEl = document.getElementById("mathExpression");
const precisionEl = document.getElementById("precision");

let isResultNegative = false;

formEl.addEventListener("submit", (e) => {
  e.preventDefault();

  if (
    !aEl.value.trim() ||
    !bEl.value.trim() ||
    !mathExpressionEl.value.trim() ||
    !precisionEl.value.trim()
  ) {
    console.error("all fields should be filled");
    return;
  }
  const params = {
    a: +aEl.value,
    b: +bEl.value,
    mathExpression: mathExpressionEl.value,
    precision: +precisionEl.value,
  };

  editParams(params);

  getResult(params);
});

async function getResult(params) {
  toggleLoader();

  const response = await fetch("/api/calculate-integral", {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify(params),
  });
  let { data } = await response.json();
  data = Number(data);
  if (isResultNegative) {
    data = -data;
  }
  resultEl.value = data.toFixed(
    params.precision.toString().split(".")[1]?.length || 1
  );
  resultEl.value = formatResult(data, params.precision);

  toggleLoader();
}

function editParams(params) {
  if (params.a > params.b) {
    let temp = params.a;
    params.a = params.b;
    params.b = temp;
    isResultNegative = true;
  } else {
    isResultNegative = false;
  }
  params.mathExpression = params.mathExpression
    .replaceAll("^", "**")
    .toLowerCase();
}

function toggleLoader() {
  loaderEl.classList.toggle("invisible");
  resultGroupEl.classList.toggle("invisible");
}

function formatResult(data, precision) {
  const intLength = data.toString().split(".")[0].length;

  if (intLength > 10) {
    return `${(data / 10 ** (intLength - 1)).toFixed(2)} * 10^${intLength - 1}`;
  }

  return data.toFixed(precision.toString().split(".")[1]?.length || 1);
}
