namespace TodolistApi.Models;

public class Todo
{
    public int Id { get; set; }
    public TodoType Type { get; set; }
    public string Title { get; set; }
    public string? Description { get; set; }
    public DateTime DueDate { get; set; }
    public DateTime? CompletedAt { get; set; }
    public List<User> Assignees { get; set; }
}
