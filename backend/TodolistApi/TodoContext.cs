using Microsoft.EntityFrameworkCore;
using TodolistApi.Models;

namespace TodolistApi;

public class TodoContext : DbContext
{
    public DbSet<User>? Users { get; set; }
    public DbSet<Todo>? Todos { get; set; }
    public DbSet<TodoType>? TodoTypes { get; set; }

    private string DbPath { get; set; }

    public TodoContext()
    {
        const Environment.SpecialFolder folder = Environment.SpecialFolder.LocalApplicationData;
        DbPath = $"{Environment.GetFolderPath(folder)}{Path.DirectorySeparatorChar}todo.db";
    }

    protected override void OnConfiguring(DbContextOptionsBuilder options)
        => options.UseSqlite($"Data Source={DbPath}");
}
