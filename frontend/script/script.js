const form = document.getElementById("criar-tarefa");
const botaoDeletar = document.getElementById("botao-deletar")

// Ao recarregar a página as tarefas já irão aparecer
carregarTarefas()


// Função par aparecer as tarefas na tela
async function carregarTarefas() {
    const response = await fetch("http://localhost:3000/listar");
    const tarefas = await response.json();

    const listar = document.getElementById("lista-tarefas");
    listar.innerHTML = ""; // Limpa a lista atual

    tarefas.forEach(tarefa => {
        const li = document.createElement("li");
        li.textContent = `${tarefa.Descricao} - ${tarefa.Feita ? "" : "Pendente"}`;

        // Construção da função de deletar tarefa: Início
        const botaoDeletar = document.createElement("button")
        botaoDeletar.textContent = "Deletar"
        botaoDeletar.addEventListener("click", async () => {
            await fetch(`http://localhost:3000/deletar?id=${tarefa.Id}`, {
                method: "DELETE",
            });
            carregarTarefas(); // Recarrega a lista após deletar
        });
        
        // Construção da função de deletar tarefa: Fim

        // Construção do botão de edição da tarefa: Ínicio
        const botaoEditar = document.createElement("button");
        botaoEditar.textContent = "Editar";

        botaoEditar.addEventListener("click", async () => {
            const novoTexto = prompt("Digite a nova descrição:", tarefa.Descricao);

            if (novoTexto && novoTexto.trim() !== "") {
                const response = await fetch(`http://localhost:3000/editar?id=${tarefa.Id}&descricao=${encodeURIComponent(novoTexto)}`, {
                    method: "PUT",
                });

                if (response.ok) {
                    console.log("Tarefa atualizada com sucesso");
                    carregarTarefas();
                } else {
                    console.error("Erro ao atualizar tarefa");
                }
            }
        });

        // Contrução do botão de edição da tarefa: Fim

        li.appendChild(botaoEditar);

        li.appendChild(botaoDeletar)

        listar.appendChild(li);
    });
}

// Criação de tarefas
form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const descricao = document.getElementById("descricao").value;

    const response = await fetch("http://localhost:3000/criar", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            Descricao: descricao
        }),
    });

    const data = await response.json();
    console.log("Resposta do servidor ao criar tarefa:", data); // Adicionado para depuração
    carregarTarefas(); // Atualiza a lista de tarefas
});
