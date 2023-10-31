<?php

use App\Http\Controllers\Auth\AllUsers;
use App\Http\Controllers\Auth\Login;
use App\Http\Controllers\Auth\Me;
use App\Http\Controllers\Task\AllTasks;
use App\Http\Controllers\Task\CreateTask;
use App\Http\Controllers\Task\DeleteTask;
use App\Http\Controllers\Task\UpdateTask;
use App\Http\Controllers\TaskType\AllTaskTypes;
use App\Http\Controllers\TaskType\CreateTaskType;
use App\Http\Controllers\TaskType\DeleteTaskType;
use App\Http\Controllers\TaskType\UpdateTaskType;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::prefix('auth')->group(function () {
    Route::get('me', Me::class);
    Route::post('login', Login::class);
    Route::get('users', AllUsers::class);
});

Route::prefix('tasks')->group(function () {
    Route::get('all', AllTasks::class);
    Route::post('create', CreateTask::class);
    Route::patch('update/{task}', UpdateTask::class);
    Route::delete('delete/{task}', DeleteTask::class);
});

Route::prefix('taskTypes')->group(function () {
    Route::get('all', AllTaskTypes::class);
    Route::post('create', CreateTaskType::class);
    Route::patch('update/{taskType}', UpdateTaskType::class);
    Route::delete('delete/{taskType}', DeleteTaskType::class);
});
