Manowar DevOps Project

Este projeto demonstra um fluxo completo de CI/CD com GitOps para aplicações containerizadas em Kubernetes:

• 	Aplicação e pipeline de build/testes: https://github.com/vi-tineo/ManowarImage-CI.git
• 	Repositório GitOps monitorado pelo Argo CD: https://github.com/vi-tineo/ManowarK8sGitOps-CI.git


Fluxo

• 	CI compila o código-fonte, executa testes unitários e gera a imagem Docker
• 	A imagem é publicada no Docker Hub com tag baseada no commit
• 	O workflow atualiza o repositório GitOps com a nova versão
• 	O Argo CD detecta a mudança e aplica no cluster bare metal


Tecnologias

• 	GitHub Actions
• 	Docker & DockerHub
• 	Kubernetes
• 	Argo CD
• 	GitOps
• 	Go


Fluxo: [App Repo] → [CI Pipeline] → [Docker Hub] ↓ [GitOps Repo] → [Argo CD] → [Cluster]


Nota Importante

Este portfólio apresenta um fluxo simplificado de CI/CD e GitOps para fins de demonstração. Elementos essenciais em ambientes de produção, como segurança segregada (RBAC, políticas de rede, gestão de segredos), observabilidade (monitoramento, métricas, tracing e logging estruturado) e Service Mesh (para controle de tráfego, segurança e visibilidade entre serviços), não foram incluídos aqui para manter o foco na automação de deploys.
Em implementações reais, tais componentes seriam incorporados juntamente com infraestrutura como código (IaC), garantindo confiabilidade, segurança e visibilidade operacional em escala.
