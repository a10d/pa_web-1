<?php

namespace Database\Factories;

use App\Models\Task;
use App\Models\TaskType;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends Factory<Task>
 */
class TaskFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'task_type_id' => TaskType::factory(),
            'title' => $this->faker->word,
            'description' => $this->faker->paragraph,
            'dueDate' => $this->faker->dateTimeBetween('now', '+1 month'),
            'completed' => $this->faker->boolean,
        ];
    }

    public function withRandomAssigneesFrom(Collection $users)
    {
        return $this->afterCreating(function (Task $task) use ($users) {
            $task
                ->taskAssignees()
                ->attach(
                    $users->random(rand(1, 3))->pluck('id')
                );
        });
    }

    public function withRandomType(Collection $taskTypes)
    {
        return $this->state(function () use ($taskTypes) {
            return [
                'task_type_id' => $taskTypes->random()->id,
            ];
        });
    }
}
