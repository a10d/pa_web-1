namespace TodolistApi.Models;

public class TodoType
{
    public int Id { get; set; }
    public string Name { get; set; }
    public string? Description { get; set; }
    public string Color { get; set; }
    public int ReminderTime { get; set; }
}
