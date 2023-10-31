FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY ["TodolistApi/TodolistApi.csproj", "TodolistApi/"]
RUN dotnet restore "TodolistApi/TodolistApi.csproj"
COPY . .
WORKDIR "/src/TodolistApi"
RUN dotnet build "TodolistApi.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "TodolistApi.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "TodolistApi.dll"]
