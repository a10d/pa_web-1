<?php

namespace Tests\Feature;

use App\Models\Task;
use App\Models\TaskType;
use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Http\Response;
use Illuminate\Support\Arr;
use Tests\TestCase;

class TaskTest extends TestCase
{
    use RefreshDatabase;

    public function test_a_guest_cannot_list_any_tasks()
    {
        $this->get('/api/tasks/all')->assertStatus(Response::HTTP_FORBIDDEN);
    }

    public function test_any_user_can_list_all_tasks()
    {
        $this->actingAs(User::factory()->create());

        Task::factory()->count(5)->create();

        $this->get('/api/tasks/all')
            ->assertOk()
            ->assertJsonCount(5, 'tasks');
    }

    public function test_a_user_can_create_a_task()
    {
        $data = [
            'type' => TaskType::factory()->create()->id,
            'title' => 'Test Task',
            'description' => 'Test Task Description',
            'dueDate' => now()->addDay()->format('Y-m-d H:i:s'),
        ];

        $this
            ->actingAs(User::factory()->create())
            ->post('/api/tasks/create', $data)
            ->assertOk();

        $this->assertDatabaseHas('tasks', [
            'task_type_id' => $data['type'],
            ...Arr::except($data, ['type'])
        ]);
    }

    public function test_a_user_can_update_a_task()
    {
        $task = Task::factory()->create([
            'title' => 'Task before update',
        ]);

        $this->assertEquals('Task before update', $task->title);

        $this
            ->actingAs(User::factory()->create())
            ->patch('/api/tasks/update/' . $task->id, [
                ...Arr::only($task->toArray(), ['type', 'title', 'description', 'dueDate', 'completed']),
                'type' => $task->taskType->id,
                'title' => 'Task after update',
            ])
            ->assertOk();

        $task->refresh();

        $this->assertEquals('Task after update', $task->title);
    }

    public function test_a_user_can_delete_an_existing_task()
    {
        $task = Task::factory()->create();

        $this
            ->actingAs(User::factory()->create())
            ->delete('/api/tasks/delete/' . $task->id)
            ->assertOk();

        $this->assertDatabaseMissing('tasks', [
            'id' => $task->id,
        ]);
    }
}
