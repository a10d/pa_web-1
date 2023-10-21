<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use App\Models\Task;
use App\Models\TaskType;
use App\Models\User;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        $users = User::factory()->createMany([
            [
                'username' => 'admin',
                'password' => 'admin',
            ],
            [
                'username' => 'user1',
                'password' => 'user1',
            ],
            [
                'username' => 'user2',
                'password' => 'user2',
            ]
        ]);

        $types = TaskType::factory()
            ->count(5)
            ->create();

        Task::factory()->count(50)
            ->withRandomAssigneesFrom($users)
            ->withRandomType($types)
            ->create();
    }
}
