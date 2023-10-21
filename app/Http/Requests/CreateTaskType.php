<?php

namespace App\Http\Requests;

use App\Models\TaskType;
use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Support\Facades\Auth;

class CreateTaskType extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return Auth::check()
            && Auth::user()->can('create', [TaskType::class]);
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'name' => ['string', 'required'],
            'description' => ['string', 'required'],
            'color' => ['string', 'required', 'regex:/^#[0-9a-f]{6}$/i'],
            'reminderTime' => ['integer', 'required', 'min:1', 'max:60'],
        ];
    }
}
